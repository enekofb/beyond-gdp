FROM golang:1.17-alpine AS build
RUN apk add gcc libc-dev
WORKDIR /go/src/app

COPY . .
COPY .resources/prod/world-happiness-data.csv /world-happiness-data.csv
COPY .resources/prod/config.yaml /config.yaml
RUN go test  ./...
RUN go build -o /bin/world-happiness-api



FROM alpine:3.11
RUN apk add ca-certificates

COPY --from=build /bin/world-happiness-api /bin/world-happiness-api
COPY --from=build /world-happiness-data.csv /world-happiness-data.csv
COPY --from=build /config.yaml /config.yaml
CMD ["world-happiness-api"]
