FROM golang:1.17-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY . .
RUN go test  ./...
RUN go build -o /bin/world-happiness-api


FROM alpine:3.11
RUN apk add --no-cache ca-certificates

COPY --from=build /bin/world-happiness-api /bin/world-happiness-api
CMD ["world-happiness-api"]
