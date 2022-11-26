package main

import (
    "context"
	"strings"
	"time"
	"fmt"
	"net/http"		
	"log"	
	"bytes"
	"io/ioutil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/julienschmidt/httprouter"
	"github.com/segmentio/encoding/json"	
	//"os"
	//"github.com/gorilla/mux"
	//"github.com/arieldll/free5gc-ariel/blob/main/NFs/nef/model_nef_event_exposure_subsc"
	"github.com/google/uuid"
    "github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/models"	
	"github.com/urfave/cli"
)

const DataCollectionNefRegistration = "datacollection.nef2.Registration"

type RegistrationObject struct {
	Id string
	Addr string
	Type string
}

type SubRequest struct {
	Id string `json:"id"`
	Addr string `json:"addr"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type SendDataStruct struct { 
	Data string `json:"data"`
}

func GetMongoDBUri()string{
	//escrever a leitura posterior
	return "mongodb://127.0.0.1:27017"
	//return GetConfiguration().MongoURI
}

func GetDBName()string{
	//ler do yaml
	return "free5gc"
}

func CloseConnection(cli*mongo.Client, ctx context.Context){
	cli.Disconnect(ctx)
}

func GetMongoConnection() (*mongo.Database, *mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBUri()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	conn := client.Database(GetDBName())
	return conn, client, ctx
}

func GetCollectionsName()[]string {
	var CollectionNames []string
	db, client, ctx := GetMongoConnection()
	result, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, coll := range result {
		CollectionNames = append(CollectionNames, coll)
	}
	CloseConnection(client, ctx)
	return CollectionNames
}

func AddRegistrationAccept(data *RegistrationObject){	
	db, client, ctx  := GetMongoConnection()
	collection := db.Collection(DataCollectionNefRegistration)
	_, err := collection.InsertMany(context.TODO(), []interface{}{data})

	fmt.Println("The Registration was saved on the database...")
	/* close MONGO connection */
	CloseConnection(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveAFsRegistered(id string){
	db, client, ctx  := GetMongoConnection()
	collection := db.Collection(DataCollectionNefRegistration)
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err !=nil {
        log.Fatal(err)
    }

	for cur.Next(context.TODO()) {
		var elem RegistrationObject
        err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		if(id == elem.Id){
			result, err := collection.DeleteOne(context.TODO(), elem)
			if err != nil {
				log.Fatal(err)
			}			
			fmt.Printf("The Registration removed %v AF: %v\n", result.DeletedCount, id)
			break
		}				
	}
	CloseConnection(client, ctx)
	/*result, err := podcastsCollection.DeleteOne(ctx, bson.M{"title": "The Polyglot Developer Podcast"})
	if err != nil {
		log.Fatal(err)
	}*/
}

func GetAFsRegistered(afType string)[] RegistrationObject{
	db, client, ctx := GetMongoConnection()
	collection := db.Collection(DataCollectionNefRegistration)
	findOptions := options.Find()
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err !=nil {
        log.Fatal(err)
    }
	var results[] RegistrationObject
	for cur.Next(context.TODO()) {
		var elem RegistrationObject
        err := cur.Decode(&elem)
		if err != nil {
            log.Fatal(err)
        }
		if elem.Type == afType {
			results = append(results, elem)
		}
		//fmt.Println(elem.Id)
	}
	//fmt.Println(collection)	
	//fmt.Println(client, ctx)
	//_, err := collection.InsertMany(context.TODO(), []interface{}{data})

	/* close MONGO connection */
	CloseConnection(client, ctx)
	/*if err != nil {
		log.Fatal(err)
	}*/
	return results
}

func goLinkById(s string, data string){
	client := &http.Client{}
	//get link by id
	l := s //"https://back.placafipe.xyz/solicitacoes/total-por-dia-usuario"
	
	var sendData SendDataStruct;
	sendData.Data = data
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(sendData)
	req, err := http.NewRequest("POST", l, &buff)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Println(req)
    
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }	
	fmt.Println(resp)

	defer resp.Body.Close()

    // Read Response Body
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

	//return resp
	// Display Results
    //fmt.Println("response Status : ", resp.Status)
    //fmt.Println("response Headers : ", resp.Header)
    fmt.Println("response Body : ", string(respBody))
}

func fireHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	start := time.Now()
	//id := ps.ByName("id")
	//fmt.Println("id ", id)
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(r.Body)
	var post SubRequest 
	json.Unmarshal(reqBody, &post)
	fireType := post.Type
	dataReceived := post.Data
	registers := GetAFsRegistered(fireType)
	elapsed := time.Since(start)
	//log.Printf("Time for execution %s", elapsed)
	fmt.Println("Time for execution", len(registers), ";", elapsed, dataReceived)
	//fmt.Println(" ---- Fire!! ---- ", id)

	for i, v := range registers{
		fmt.Println(i)
		fmt.Println(v.Addr)
		goLinkById(v.Addr, dataReceived)
	}
	
	//rec := goLinkById(id)
	//fmt.Println("rec ", rec)
}

func unsubscriptionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params){	
	id := ps.ByName("id")
	fmt.Println("parametro ", id)
	RemoveAFsRegistered(id)
	w.WriteHeader(http.StatusNoContent)
}

func updateSubscriptionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params){	
	id := ps.ByName("id")
	fmt.Println("parametro ", id)
	RemoveAFsRegistered(id)
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.Body)
	var post SubRequest 
	json.Unmarshal(reqBody, &post)		

	var oRegister RegistrationObject;
	oRegister.Addr = post.Addr
	oRegister.Id = post.Id
	oRegister.Type = post.Type
	AddRegistrationAccept(&oRegister)
}

//func subscriptionHandler(w http.ResponseWriter, r *http.Request) {	
func subscriptionHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {	
    if r.Method == "POST" && r.URL.Path == "/subscriptions" {
		//var p SubRequest
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println(r.Body)
		var post SubRequest 
		json.Unmarshal(reqBody, &post)		

		var oRegister RegistrationObject;
		oRegister.Addr = post.Addr
		oRegister.Id = post.Id
		oRegister.Type = post.Type
		AddRegistrationAccept(&oRegister)

		fmt.Println(oRegister.Id, oRegister.Addr, oRegister.Type)
		
		w.WriteHeader(http.StatusCreated)
		return;
		//fmt.Fprintf(w, "BODYYYY", r.Body)
		//reqBody, _ := ioutil.ReadAll(r.Body)
		//json.Unmarshal(reqBody, &p)
		/*err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			fmt.Fprintf(w, ">>> ERRO <<<")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}*/
        //fmt.Fprintf(w, "id/addr:", p.id, " -- ", p.addr)
		//fmt.write
		//return
    }

	http.Error(w, "404 not found.", http.StatusNotFound)
    return    
}


func main() {
	app := cli.NewApp()
	app.Name = "nef"
	app.Usage = "5G NEF"	
	
	//endereço/porta do NRF
    nrfUri := "http://127.0.0.10:8000"
    
	//arquivo de configuração
    configuration := Nnrf_NFManagement.NewConfiguration()
	configuration.SetBasePath(nrfUri)
	client := Nnrf_NFManagement.NewAPIClient(configuration)	
    nfInstanceId := uuid.New().String()
	
	//criar um novo profile para a função
    var profile models.NfProfile	
	profile.NfInstanceId = nfInstanceId
	profile.NfType = models.NfType_NEF
	profile.NfStatus = models.NfStatus_REGISTERED

	//ip, porta e contexto (http/https) de registro da nova NF (neste caso, NRF)
	register_ipv4 := "127.0.0.1"
	sbi_port := 29895
	context_urischeme := models.UriScheme_HTTP

	apiPrefix := fmt.Sprintf("%s://%s:%d", context_urischeme, register_ipv4, sbi_port)
	
    fmt.Println("Trying to connect to the NRF", nfInstanceId, " Register", apiPrefix)	
    for {
		//enviar o registro da nova NF
		_, res, err := client.NFInstanceIDDocumentApi.RegisterNFInstance(context.TODO(), nfInstanceId, profile)
		if err != nil || res == nil {			
			fmt.Println(fmt.Errorf("NEF register to NRF Error[%s]", err.Error()))
			time.Sleep(2 * time.Second)
			continue
		}		
		status := res.StatusCode
		fmt.Println("Received Status: ", status)
		if status == http.StatusOK {
            fmt.Println("Status - OK = Function Already Registered")
			break
		} else if status == http.StatusCreated {
            fmt.Println("Status - Created and Registered")			
			resourceUri := res.Header.Get("Location")			
			retrieveNfInstanceId := resourceUri[strings.LastIndex(resourceUri, "/")+1:]
			fmt.Println(resourceUri, retrieveNfInstanceId, err)
			break
		} else {
			fmt.Println("handler returned wrong status code", status)
			fmt.Println("NRF return wrong status code", status)
			break
		}
	}

	//
	//dkt := GetCollectionsName()
	//fmt.Println(dkt)

	//GetAFsRegistered("type1")

	router := httprouter.New()
	//http.HandleFunc("/subscriptions", subscriptionHandler) // Update this line of code	
	router.POST("/subscriptions", subscriptionHandler)
	router.PUT("/subscriptions/:id", updateSubscriptionHandler)
	router.DELETE("/subscriptions/:id", unsubscriptionHandler)
	router.POST("/fire/:id", fireHandler)
	fmt.Printf("Starting server at port 20000\n")

	//goLinkById("1")
    if err := http.ListenAndServe(":20000", router); err != nil {
        log.Fatal(err)
    }
}