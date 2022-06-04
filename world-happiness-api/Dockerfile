FROM golang:1.17-alpine AS build
RUN apk add gcc libc-dev
WORKDIR /go/src/app

COPY . .
COPY .resources/world-happiness-data.csv /world-happiness-data.csv
RUN go test  ./...
RUN go build -o /bin/world-happiness-api



FROM alpine:3.11
RUN apk add ca-certificates

COPY --from=build /bin/world-happiness-api /bin/world-happiness-api
COPY --from=build /world-happiness-data.csv /world-happiness-data.csv
CMD ["world-happiness-api"]
