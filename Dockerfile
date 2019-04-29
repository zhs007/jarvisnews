
FROM golang:1.12 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR /src/jarvisnews

COPY ./go.* /src/jarvisnews/

RUN go mod download

COPY . /src/jarvisnews

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvisnews . \
    && mkdir /app \
    && mkdir /app/jarvisnews \
    && mkdir /app/jarvisnews/dat \
    && mkdir /app/jarvisnews/logs \
    && mkdir /app/jarvisnews/cfg \
    && cp ./jarvisnews /app/jarvisnews/ \
    && cp cfg/config.yaml.default /app/jarvisnews/cfg/config.yaml

FROM alpine
WORKDIR /app/jarvisnews
COPY --from=builder /app/jarvisnews /app/jarvisnews
CMD ["./jarvisnews", "start"]
