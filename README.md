# trabalho-cmu
Trabalho de computação móvel e ubíqua

Este trabalho implementa uma NEF (Network Exposure Function) conceitual para o projeto free5GC, com todos os métodos necessários a exposição de eventos através das definições da 3GPP para Nnef_EventExposure

## Requisitos: 
- Core funcional do projeto free5GC
- Linguagem go versão 1.17 ou superior 
- Bibliotecas da linguagem go 

## Como instalar o go? 
https://go.dev/dl/

## Biliotecas necessárias
context, strings, time, fmt, net/http, log, bytes, ioutil, mongodb, httprouter, json, uuid, openapi, cli

## Instalando biliotecas de terceiros no go:
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

Faça um clone do projeto oficial: 
'''git clone https://github.com/free5gc/free5gc'''

Configure o core do free5GC de modo que a função NRF esteja exposta e funcionando. Você pode ter mais suporte com isso <a href="https://github.com/free5gc/free5gc/wiki">aqui</a>

Após terminado o clone, entre na pasta free5gc/NFs e clone este repositório:  

```
$ git clone https://github.com/arieldll/trabalho-cmu
```

No arquivo nef.go, configure o seguinte: 

Na linha 43, insira o endereço e porta do MongoDB do projeto free5gc

Na linha 252, configure o endereço e porta da NRF

Após, apenas execute: 

```
$ go run nef.go
```

Se tudo estiver funcionando, a NEF estará disponível e exposta na porta 20000





