#syntax
FROM golang:1.16-alpine
WORKDIR /controller
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o math .
CMD [ "./math" ]