FROM golang:1.11.0-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get -u github.com/gorilla/mux && \
    go get -u github.com/gorilla/rpc && \
    go get -u github.com/micro/protobuf/protoc-gen-go && \
    go get -u github.com/mchirico/gog/cmd/gog  && \
    go get -u github.com/mchirico/gog/cmd/grpc
                    
# ENV SOURCES /go/src/github.com/mchirico/gopb/Communication/Go-Micro/
# COPY . ${SOURCES}

# RUN cd ${SOURCES}server/ && CGO_ENABLED=0 go build

# ENV CONSUL_HTTP_ADDR localhost:8500

# WORKDIR ${SOURCES}server/
# CMD ${SOURCES}server/server
CMD gog

# To build this:
# docker build -t mchirico/gog:latest .


# To run this: 
# docker run -p 8080:8080 --rm -it mchirico/gog:latest


# To push
# docker push mchirico/gog:latest
