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
### Set up database
Start cockroach sql command line :
```
cockroach sql --insecure
```

Then excecute code on __SQL/kontak.sql__ to cockroach sql command line.

### Run cockroach
Run cockroach on insecure mode :
```
cockroach start --insecure
```

### Run golang app
Go to app directory, then run golang app :
```
go run main.go
```

### Open app on browser
Open your browser, then access :
```
localhost:8010
```
