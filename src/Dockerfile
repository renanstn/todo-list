FROM golang:1.22.1

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . .

CMD [ "go", "run", "." ]
