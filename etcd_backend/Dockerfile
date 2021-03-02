FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/etcd_backend/
RUN cd /go/etcd_backend && CGO_ENABLED=0 go build

FROM alpine:3.13.2
EXPOSE 8082
WORKDIR /app
COPY --from=gobuilder /go/etcd_backend/etcd_backend.yaml /app/etcd_backend.yaml
COPY --from=gobuilder /go/etcd_backend/etcd_backend /app/etcd_backend
CMD ["./etcd_backend"]