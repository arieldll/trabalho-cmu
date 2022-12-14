<p align="left"> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.linux.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="40"/> </a> <a href="https://www.mongodb.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mongodb/mongodb-original-wordmark.svg" alt="mongodb" width="40" height="40"/> </a> </p>

# Network Exposure Function (Português)

Desenvolvimento de Software Orientado à Computação Móvel e Ubíqua

Este trabalho implementa uma Network Exposure Function (NEF) conceitual para o projeto <a href="https://github.com/free5gc/free5gc">free5gc</a>, com todos os endpoints definidos pela exposição de eventos, através das definições da 3GPP para <a href="https://github.com/jdegre/5GC_APIs/blob/Rel-18/TS29591_Nnef_EventExposure.yaml">TS29591_Nnef_EventExposure</a>. Funciona em conjunto com a Application Function (AF) disponível <a href="https://github.com/opoze/5gaf">aqui</a>.

## Arquitetura
A arquitetura de funcionamento segue o seguinte fluxo: 
![image](https://user-images.githubusercontent.com/2493503/204649619-ca5960f9-1106-4c24-b32a-4c3f5193f4ed.png)


## Requisitos de software
- Core funcional do projeto <a href="https://github.com/free5gc/free5gc">free5gc</a>
- Linguagem go versão 1.17 ou superior 
- Ubuntu 22.04 ou superior 
- Bibliotecas da linguagem go 

## Requisitos de hardware 
- Processador com no mínimo 2 GHz 
- 4 Gigabytes de memória RAM 
- 2 Gigabyte de espaço em disco livre 
- Sistema Operacional Linux

## Como instalar o go? 
https://go.dev/dl/

## Biliotecas necessárias
context, strings, time, fmt, net/http, log, bytes, ioutil, mongodb, httprouter, json, uuid, openapi, cli

## Instalando biliotecas de terceiros no go
```
$ go get go.mongodb.org/mongo-driver/bson
$ go get go.mongodb.org/mongo-driver/mongo
$ go get go.mongodb.org/mongo-driver/mongo/options
$ go get github.com/julienschmidt/httprouter
$ go get github.com/segmentio/encoding/json
$ go get github.com/google/uuid
$ go get github.com/free5gc/openapi/Nnrf_NFManagement
$ go get github.com/free5gc/openapi/models
$ go get github.com/urfave/cli
```


## Por onde começar

Faça um clone ou fork do projeto oficial do free5gc:

```
git clone https://github.com/free5gc/free5gc
```

Instale todas as dependências e compile as funções disponíveis no projeto, <a href="https://github.com/free5gc/free5gc">free5gc</a>, <a href="https://github.com/free5gc/free5gc/wiki/Installation">conforme este link</a>. Para que a NEF funcione, faz-se necessário apenas a compilação e execução da Network Repository Function (NRF), junto com o MongoDB. Fica a seu critério compilar ou não as outras funções do core.  

Feito a instalação e compilação da NRF, inicialize-a. Se você estiver usando o projeto free5gc, a partir da raiz do projeto

```
$ ./bin/nrf 
```

Tendo a NRF instanciada, entre na pasta free5gc/NFs e clone este próprio repositório: 

```
$ git clone https://github.com/arieldll/trabalho-cmu
```

No arquivo nef.go, configure o seguinte: 

```
Na linha 43, insira o endereço e porta do MongoDB do projeto free5gc
```

```
Na linha 252, configure o endereço e porta da NRF
```

Após, apenas execute: 

```
$ go run nef.go
```

Se tudo estiver funcionando, a NEF estará disponível e exposta na porta 20000

## Primeiros resultados
A seguir, as primeiras avaliações de desempenho da NEF conceitual, utilizando uma quantidade incremental de AFs registradas, incrementando o registor de AFs ao longo do tempo. Todos os valores foram coletados via instrumentação da aplicação, e estão descritos na ordem de milissegundos (ms). Os experimentos foram feitos com 1, 5, 10, 50 e 100 AFs e uma instância da NEF. As requisições de inscrição foram feitas através do <a href="https://www.postman.com/">Postman</a>. 
O computador host do core do free5gc e NEF tem a seguinte configuração: 
- Processador Intel(R) Xeon(R) CPU E5-2650 0 @ 2.00GHz – 1 core 
- 4GB de memória RAM 


![image](https://user-images.githubusercontent.com/2493503/205521049-edbca8f3-f395-427a-87cb-00625b21d2f7.png)

Podemos notar que o tempo de resposta é proporcional a quantidade de AFs registradas. 

![image](https://user-images.githubusercontent.com/2493503/204617391-c5e92409-1b02-4c3f-87f7-2d58bf6a8918.png)


# Network Exposure Function (English)

This work implements a conceptual Network Exposure Function (NEF) to the <a href="https://github.com/free5gc/free5gc">free5gc project</a>, covering all endpoints defined by the event exposure through 3GPP definitions for the <a href="https://github.com/jdegre/5GC_APIs/blob/Rel-18/TS29591_Nnef_EventExposure.yaml">TS29591_Nnef_EventExposure</a>. This project works together the Application Function (AF) available <a href="https://github.com/opoze/5gaf">here</a>.

## Architecture
The architecture works as flows: 
![image](https://user-images.githubusercontent.com/2493503/205521236-b03eeff8-6a9f-40ad-a259-7832be74ca52.png)


## Software requirements
- Functional core of <a href="https://github.com/free5gc/free5gc">free5gc</a> project 
- Go language version 1.17 or higher
- Ubuntu 22.04 or higher  
- Go language 3rd party libraries

## Hardware requirements
- Processor minimum 2 GHz 
- 4 Gigabytes of RAM
- 2 Gigabytes of free disk space
- Linux Operating System

## How to install golang? 
https://go.dev/dl/

## 3rd party libraries of golang
context, strings, time, fmt, net/http, log, bytes, ioutil, mongodb, httprouter, json, uuid, openapi, cli

## Installing 3rd party libraries in golang
```
$ go get go.mongodb.org/mongo-driver/bson
$ go get go.mongodb.org/mongo-driver/mongo
$ go get go.mongodb.org/mongo-driver/mongo/options
$ go get github.com/julienschmidt/httprouter
$ go get github.com/segmentio/encoding/json
$ go get github.com/google/uuid
$ go get github.com/free5gc/openapi/Nnrf_NFManagement
$ go get github.com/free5gc/openapi/models
$ go get github.com/urfave/cli
```

## How to start

Clone or fork the official free5gc project: 

```
git clone https://github.com/free5gc/free5gc
```

Install all dependencies and compile the NFs available in the project, <a href="https://github.com/free5gc/free5gc">free5gc</a>, as such<a href="https://github.com/free5gc/free5gc/wiki/Installation">this link</a>. For NEF works, you need to compile and execute Network Repository Function (NRF), together with MongoDB. If necessary, you would compile the other 5G core functions.  

After, start NRF. From 5gc root project: 

```
$ ./bin/nrf 
```
Enter on the free5gc/NFs folder and clone this one repository: 

```
$ git clone https://github.com/arieldll/trabalho-cmu
```

On the nef.go file, configure as follow: 

```
In line 43, change to the IP/port of MongoDB
```

```
In line 252, change to the IP/port of NRF
```

If all dependencies are correct, just execute: 

```
$ go run nef.go
```

If all is working, NEF is available at http://localhost:20000

## First results
Above, are the first performance evaluations of this implementation of conceptual NEF. For this evaluation, are used an incremental amount of registered AFs, scaling the register of AFs over time. All values were collected via the application's instrumentation, and all values are described in the order of milliseconds (ms). The experiments were carried out with 1, 5, 10, 50, and 100 AFs and just one instance of NEF. The subscr
tion requests were made using <a href="https://www.postman.com/">Postman</a>.
The free5gc and NEF core host computer has the following configuration:
- Intel(R) Xeon(R) processor E5-2650 CPU 0 @ 2.00GHz – 1 core
- 4GB of RAM


![image](https://user-images.githubusercontent.com/2493503/205520904-2e901179-0c44-46ff-b6d0-92cfbe04d511.png)

We can see a response time is proportional to the amount of registered AFs

![image](https://user-images.githubusercontent.com/2493503/205520824-4152e018-a216-4dec-9d24-2482f2ab13a8.png)
