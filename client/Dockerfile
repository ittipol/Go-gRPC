FROM golang:1.19  as build

WORKDIR /go-app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./app

# =================================================

FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/app ./
COPY --from=build /go-app/config.yaml ./

ENV TLS_CERT=/app/tls/ca.crt

# ENTRYPOINT ["./app"]
CMD [ "tail", "-F", "anyfile" ]