package main

import (
    //"context"
	"fmt"
	"net/http"
	//"strings"
	//"time"
	"github.com/segmentio/encoding/json"
	"log"	
	"io/ioutil"
	//"github.com/gorilla/mux"
	//"github.com/arieldll/free5gc-ariel/blob/main/NFs/nef/model_nef_event_exposure_subsc"
	//"github.com/google/uuid"
    //"github.com/free5gc/openapi/Nnrf_NFManagement"
	//"github.com/free5gc/openapi/models"	
	//"github.com/urfave/cli"
)

type SubRequest struct {
	Id string `json:"id"`
	Addr string `json:"addr"`
}

func unsubscriptionHandler(w http.ResponseWriter, r *http.Request) {	

}

func subscriptionHandler(w http.ResponseWriter, r *http.Request) {	
    if r.Method == "POST" && r.URL.Path == "/subscriptions" {
		//var p SubRequest
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Println(r.Body)
		var post SubRequest 
		json.Unmarshal(reqBody, &post)

		fid := post.Id
		faddr := post.Addr

		fmt.Println(fid, faddr)
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
	/*app := cli.NewApp()
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
	}*/
	
	http.HandleFunc("/subscriptions", subscriptionHandler) // Update this line of code	
	fmt.Printf("Starting server at port 20000\n")

    if err := http.ListenAndServe(":20000", nil); err != nil {
        log.Fatal(err)
    }
}