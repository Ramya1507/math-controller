
FROM golang:alpine As builder
WORKDIR /controller
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o math .

From alpine
WORKDIR /mathcontroller
COPY --from=builder /controller /mathcontroller
CMD [ "./math" ]