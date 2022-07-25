FROM golang:1.18-alpine as builder

WORKDIR $GOPATH/src/github.com/Saucon/go-rest-api-sample


# COPY go.mod, go.sum and download the dependencies
COPY go.* ./

ENV GO111MODULE=on
ENV GOFLAGS=-mod=mod

RUN go mod tidy
RUN go mod download
RUN go get ./...

# COPY All things inside the project and build
COPY . .
RUN go build $GOPATH/src/github.com/Saucon/go-rest-api-sample/cmd/my_sample_app
RUN ls -al

FROM alpine:latest

# Add curl
RUN apk --no-cache add curl
# SET TZ
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

COPY --from=builder /go/src/github.com/Saucon/go-rest-api-sample/my_sample_app go-rest-api-sample/
COPY --from=builder /go/src/github.com/Saucon/go-rest-api-sample/.env ./

RUN ls -al /go-rest-api-sample

EXPOSE 9992
ENTRYPOINT [ "/go-rest-api-sample/my_sample_app" ]
