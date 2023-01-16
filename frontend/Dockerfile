FROM node:latest as builder
WORKDIR /app
COPY . .

ARG WEBSITE_DOMAIN
ENV REACT_APP_WEBSITE_DOMAIN $WEBSITE_DOMAIN

RUN npm ci
RUN npm run build

FROM nginx:alpine
RUN apk --no-cache add ca-certificates openssl

RUN openssl genrsa -des3 -passout pass:x -out server.pass.key 2048 && \
    openssl rsa -passin pass:x -in server.pass.key -out /key.pem && \
    rm server.pass.key && \
    openssl req -new -key /key.pem -out server.csr \
        -subj "/C=FI/ST=Turku/L=Turku/O=JaspnasCom/OU=IT Department/CN=reaktorbirdnest.jaspnas.com" && \
    openssl x509 -req -days 100 -in server.csr -signkey /key.pem -out /cert.crt

ENV NODE_ENV production
COPY --from=builder /app/build /usr/share/nginx/html
ADD nginx.conf /etc/nginx/nginx.conf
EXPOSE 8000
EXPOSE 44300