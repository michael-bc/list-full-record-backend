FROM golang:1.20-alpine as builder

RUN apk --no-cache add ca-certificates git

WORKDIR /go/src/github.com/michael-bc/list-full-record-backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./app ./cmd/list-full-record-backend \
    && chmod +x ./app


FROM alpine
WORKDIR /
COPY --from=builder /go/src/github.com/michael-bc/list-full-record-backend .
COPY /VERSION .
EXPOSE 8080
ENTRYPOINT ["/app"]
