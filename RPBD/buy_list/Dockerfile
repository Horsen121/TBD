FROM golang:19.1.3-alpine as build
WORKDIR /go/build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bot ./

FROM alpine:3.17 as release
WORKDIR /app

COPY --from=build /go/build/bot ./

CMD ["/app/bot"]