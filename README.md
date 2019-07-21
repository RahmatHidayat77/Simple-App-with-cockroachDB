# Simple-App-with-cockroachDB

## Prerequisites
Make sure you have Golang and CockroachDB installed on your computer. See this link to install both :

Install [Golang](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04)

Install [CockroachDB](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html)

Clone repo to your GOPATH directory :
```
git clone git@github.com:RahmatHidayat77/Simple-App-with-cockroachDB.git
```

## How to run
### Run cockroach
Run cockroach on insecure mode :
```
cockroach start --insecure
```

### Set up database
Start cockroach sql command line :
```
cockroach sql --insecure
```
Then excecute code on __SQL/kontak.sql__ to cockroach sql command line.

### Run golang app
Go to app directory, then run golang app :
```
go run main.go
```

### Open app on browser
Open your browser, then access :
[localhost:8010](localhost:8080)

## Dockerize App
First, make sure you on app directory (__Simple-App-with-cockroachDB__).
Build docker image from app :
```
docker build -t go-app:0.1 .
```

Pull cockroach image from dockerhub :
```
docker pull cockroachdb/cockroach
```

Create and run cockroachdb container :
```
docker run -d --name=roach -p 26257:26257 -p 8080:8080 cockroachdb/cockroach start --insecure
```

To access cockroachDB container we must know cockroachDB IP address.
So, save cockroachDB IP address into variabel to run app container later :
```
ROACH_IP_ADDRESS=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' roach)
```
You can echo IP address like this :
```
echo $ROACH_IP_ADDRESS
```

Access cockroachDB container to excecute sql command :
```
docker exec -it roach ./cockroach sql --insecure
```
Execute code on __SQL/kontak.sql__ to cockroach sql command line.

Create and running go-app :
```
docker run -d --name=go-app -p 8010:8010 --env ROACH_HOST=$ROACH_IP_ADDRESS go-app:0.1
```

Then access : [localhost:8010](localhost:8080).

Done.
