FROM node:16.16.0 AS FRONT
WORKDIR /static
COPY ./static .
RUN npm config set registry https://registry.npmmirror.com
RUN npm install && npm run build


FROM golang:1.18.4 AS BACK
WORKDIR /go/src/threatbook-ip-web
COPY . .
RUN ./build.sh


FROM alpine:latest AS BUILDER
LABEL MAINTAINER="https://hackerchai.com"

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add curl
RUN apk add ca-certificates && update-ca-certificates

WORKDIR /app
COPY --from=BACK /go/src/threatbook-ip-web/cmd/threatbook-ip-web/threatbook-ip-web ./server
COPY --from=BACK /go/src/threatbook-ip-web/config ./config
COPY --from=FRONT /static/dist ./static/dist
EXPOSE 8080
ENTRYPOINT ["/app/server"]