# Dockerfile for lagerauth, it compiles the site the backend and gives back a image with the files on it.

# Site:
FROM node:9.5.0 AS node-img
WORKDIR /usr/src/app
COPY ./site/package*json ./
RUN npm install
COPY ./site/ .
COPY prod.env.js ./config/prod.env.js
RUN npm run build

# Lagerauth:
FROM golang:1.10.1 AS go-img
WORKDIR /go/src/lagerauth
COPY ./lagerauth/ .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -a

# Run lagerauth:
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=go-img /go/src/lagerauth/lagerauth .
COPY --from=node-img /usr/src/app/dist ./wwwroot
COPY conf.json .
ENTRYPOINT ["./lagerauth"]