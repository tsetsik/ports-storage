FROM golang:1.15-alpine as builder

# Go build env settings
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED=0

WORKDIR /app

# Cache go modules
COPY go.mod ./
RUN go mod download

# Copy project files and build 
COPY . .
RUN go build -o ./bin/ports-storage ./main.go

CMD [ "/app/bin/ports-storage" ]