FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/api
RUN chmod +xrw /go/api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/api/cmd/api /go/api/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/api /go/api
ENTRYPOINT ["/go/api"]