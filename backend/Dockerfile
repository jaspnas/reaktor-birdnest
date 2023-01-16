FROM golang:1.19 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates openssl

RUN openssl genrsa -des3 -passout pass:x -out server.pass.key 2048 && \
    openssl rsa -passin pass:x -in server.pass.key -out /key.pem && \
    rm server.pass.key && \
    openssl req -new -key /key.pem -out server.csr \
        -subj "/C=FI/ST=Turku/L=Turku/O=JaspnasCom/OU=IT Department/CN=reaktorbirdnest.jaspnas.com" && \
    openssl x509 -req -days 100 -in server.csr -signkey /key.pem -out /cert.crt

WORKDIR /root/
COPY --from=builder /app/main ./
ADD ./settings.json ./
EXPOSE 8080
EXPOSE 44310
CMD ["./main"]