# vitica

## Prerequisites

* Redis
* MySQL 

## Setup 

INSTALLATION

* Install GO 1.4+ and configure workspace [https://golang.org/doc/install]
* Install GOM `go get github.com/mattn/gom`  
* `gom install` from project root    
* `npm install -g bunyan`  

DATABASE

* Create database "vitica" on MySQL

ENVIRONMENT VARIABLES

* `export VITICA_DIR=/<PATH_TO_PROJECT_DIR> `

## Development

Start webserver:   
`gom run webserver/main.go | bunyan`


## Testing

Run unit tests: `gom test -v`  

## Build All

```
./build.sh
```
