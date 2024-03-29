FROM golang:alpine3.13 as builder

RUN apk update && apk upgrade && apk add build-base git make sed
RUN go get github.com/silenceper/gowatch

#swagger addition via swagger-ui inyection
WORKDIR /go/src/github.com/eduardohoraciosanto/bootcamp-with-gorilla/swagger
COPY ./oas/oas.yml ./swagger.yml
RUN git clone https://github.com/swagger-api/swagger-ui && \
  cp -r swagger-ui/dist/. . && rm -r swagger-ui/ && sed -i 's+https://petstore.swagger.io/v2/swagger.json+/swagger/swagger.yml+g' index.html

WORKDIR /go/src/github.com/eduardohoraciosanto/bootcamp-with-gorilla
COPY . .

RUN GIT_COMMIT=$(git rev-parse --short HEAD) && \
  go build -o service -ldflags "-X 'github.com/eduardohoraciosanto/bootcamp-with-gorilla/config.serviceVersion=$GIT_COMMIT'"

FROM alpine:3.13 

COPY --from=builder /go/src/github.com/eduardohoraciosanto/bootcamp-with-gorilla/service /
COPY --from=builder /go/src/github.com/eduardohoraciosanto/bootcamp-with-gorilla/swagger /swagger

ENTRYPOINT [ "./service" ]