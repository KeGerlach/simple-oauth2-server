FROM golang:1.24-alpine as builder

WORKDIR /go/src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-oauth2-server -v ./cmd/app

FROM alpine:latest
COPY --from=builder /go/src/simple-oauth2-server /app/simple-oauth2-server
# COPY static/ static/

EXPOSE 8080
CMD ["/app/simple-oauth2-server"]