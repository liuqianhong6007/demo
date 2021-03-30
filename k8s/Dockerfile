FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/k8s
RUN cd /go/k8s && CGO_ENABLED=0 go build

FROM alpine:3.13.2
COPY --from=gobuilder /go/k8s/k8s /app/k8s
WORKDIR /app
ENTRYPOINT ["./k8s"]