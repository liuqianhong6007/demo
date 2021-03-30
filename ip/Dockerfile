FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/ip/
RUN cd /go/ip && CGO_ENABLED=0 go build

FROM alpine:3.13.2
COPY --from=gobuilder /go/ip/ip /app/ip
WORKDIR /app
ENTRYPOINT ["./ip"]