FROM golang:1.22.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . ./ 

RUN ls -la 

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go 

EXPOSE 8080

CMD ["/app"]


