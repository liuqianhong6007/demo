FROM golang:1.15-alpine As gobuilder
ENV GOPROXY https://goproxy.cn
COPY . /go/auth/
RUN cd /go/auth && CGO_ENABLED=0 go build

FROM alpine:3.13.2
EXPOSE 8081
WORKDIR /app
COPY --from=gobuilder /go/auth/auth.yaml /app/auth.yaml
COPY --from=gobuilder /go/auth/sql/ /app/sql/
COPY --from=gobuilder /go/auth/auth /app/auth
CMD ["./auth"]