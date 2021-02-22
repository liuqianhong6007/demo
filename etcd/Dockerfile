FROM golang:1.15-alpine As gobuilder
COPY etcd_backend /go/etcd_backend
ENV GOPROXY https://goproxy.cn
RUN cd /go/etcd_backend && CGO_ENABLED=0 go build

FROM node:current-alpine3.13 AS nodebuilder
COPY etcd-front /node/etcd_front
RUN cd /node/etcd_front && npm install && npm run build

FROM alpine:3.13.2
ENV ETCD_ADDR 127.0.0.1:2379
EXPOSE 8101
WORKDIR /app
COPY --from=gobuilder /go/etcd_backend/etcd_backend /app/etcd_backend
COPY --from=gobuilder /go/etcd_backend/etcd_backend.yaml /app/etcd_backend.yaml
COPY --from=nodebuilder /node/etcd_front/dist/ /app/html/
CMD ["./etcd_backend"]

