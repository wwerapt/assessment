FROM golang:1.19-alpine as build-base

WORKDIR /app

COPY ./expense/ ./database/

COPY *.go .

COPY go.mod .

RUN go get

RUN go mod download

RUN go test -v

RUN go build -o ./build/server .


# =======================================================

FROM alpine:3.16.2

COPY --from=build-base /app/build/server /app/server

CMD ["/app/server"]