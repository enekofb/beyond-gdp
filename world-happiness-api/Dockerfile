FROM golang:1.17-alpine AS build
RUN apk add gcc libc-dev
WORKDIR /go/src/app

COPY . .
COPY .resources/prod/world-happiness-data.csv /world-happiness-data.csv
COPY .resources/prod/config.yaml /config.yaml
RUN go test  ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/world-happiness-api

FROM public.ecr.aws/lambda/go:1

COPY --from=build /bin/world-happiness-api ${LAMBDA_TASK_ROOT}
COPY --from=build /world-happiness-data.csv ${LAMBDA_TASK_ROOT}
COPY --from=build /config.yaml ${LAMBDA_TASK_ROOT}
CMD ["world-happiness-api"]
