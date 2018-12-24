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
CMD grpc

# To build this:
# docker build -t mchirico/gog:grpc .


# To run this: 
# docker run -p 8082:8082 --rm -it mchirico/gog:grpc


# To push
# docker push mchirico/gog:grpc
