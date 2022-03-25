FROM golang:1.17-alpine
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cmd/main.go
EXPOSE 3000 3000
CMD [ "./main" ]
