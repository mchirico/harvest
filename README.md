# Harvest



## Install

```bash
# Dependencies
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/rpc
go get -u github.com/levigross/grequests
go get -u github.com/google/uuid

go get -u github.com/mongodb/mongo-go-driver/mongo


# Run Different Examples
go get -u github.com/mchirico/harvest/cmd/auth

```



# Create Minimal Docker Image

This was taken from my gog project... not implemented, yet

Reference [minimal](https://github.com/mchirico/gog/tree/min_docker/docker/minimal)


## Dockerfile
```bash
FROM       scratch

MAINTAINER Mike Chirico (chico) <mchirico@gmail.com>
ADD        gog gog

```

## Commands

```bash
#!/bin/bash
env GOOS=linux GOARCH=arm go build github.com/mchirico/gog/cmd/gog
docker build --tag gogserver .
echo -e '\n\n\nIn another window execute the following command:\n\ncurl localhost:8080\n\n'
docker run -p 8080:8080 --rm gogserver

```


## Listing of Example Programs
[awesome-go](https://awesome-go.com/)