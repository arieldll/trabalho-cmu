# trabalho-cmu
Trabalho de computação móvel e ubíqua

Este trabalho implementa uma Network Exposure Function (NEF) conceitual para o projeto free5GC, com todos os métodos necessários a exposição de eventos através das definições da 3GPP para Nnef_EventExposure

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

Faça um clone ou fork do projeto oficial do free5gc: 
```
git clone https://github.com/free5gc/free5gc
```

Instale todas as dependências do projeto. Após, configure o core do free5GC de modo que a função Network Repository Function (NRF) esteja exposta e rodando. Você pode ter informações de como proceder com isso <a href="https://github.com/free5gc/free5gc/wiki">aqui</a>

Após terminado o clone e as configurações do core do free5GC, entre na pasta free5gc/NFs e clone este próprio repositório:  

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





