<p align="left"> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.linux.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="40"/> </a> <a href="https://www.mongodb.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mongodb/mongodb-original-wordmark.svg" alt="mongodb" width="40" height="40"/> </a> </p>


Desenvolvimento de Software Orientado à Computação Móvel e Ubíqua

Este trabalho implementa uma Network Exposure Function (NEF) conceitual para o projeto <a href="https://github.com/free5gc/free5gc">free5GC</a>, com todos os métodos necessários a exposição de eventos através das definições da 3GPP para <a href="https://github.com/jdegre/5GC_APIs/blob/Rel-18/TS29591_Nnef_EventExposure.yaml">TS29591_Nnef_EventExposure</a>

Este trabalho funciona em conjunto com a Application Function (AF) disponível <a href="https://github.com/opoze/5gaf">aqui</a>

## Requisitos
- Core funcional do projeto free5GC
- Linguagem go versão 1.17 ou superior 
- Ubuntu 20.04 ou superior 
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

Inicialize a NRF. Se você estiver usando o projeto free5gc, a partir da raiz do projeto
```
$ ./bin/nrf 
```

Após terminado o clone e as configurações do core do free5gc, entre na pasta free5gc/NFs e clone este próprio repositório:  

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
A seguir, as primeiras avaliações de desempenho da NEF conceitual, utilizando uma quantidade incremental de AFs (Application Functions) registradas, simulando o incremento de AFs ao longo do tempo. Todos os valores estão descritos na ordem de milissegundos (ms). Os experimentos foram feitos com 1, 5, 10, 50 e 100 AFs e uma instância da NEF. As requisições de inscrição foram feitas através do <a href="https://www.postman.com/">Postman</a>. 
O computador host do core do free5gc e NEF tem a seguinte configuração: 
- Processador Intel(R) Xeon(R) CPU E5-2650 0 @ 2.00GHz – 1 core 
- 4GB de memória RAM 


![image](https://user-images.githubusercontent.com/2493503/204617348-ca0491a4-b48e-4542-91f4-efb2edb2edf3.png)

Podemos notar que o tempo de resposta é proporcional a quantidade de AFs registradas. 

![image](https://user-images.githubusercontent.com/2493503/204617391-c5e92409-1b02-4c3f-87f7-2d58bf6a8918.png)





