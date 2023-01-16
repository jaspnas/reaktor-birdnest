# 2023 Reaktor Birdnest solution

Submitted as part of an application to the position of developer trainee at Reaktor. 

Hosted version will be up at https://reaktorbirdnest.jaspnas.com for a limited time.

### Running the project 

get certificates from CA (e.g. letsencrypt) or run the following command 
`openssl genrsa -des3 -passout pass:x -out server.pass.key 2048 && openssl rsa -passin pass:x -in server.pass.key -out key.pem && rm server.pass.key && openssl req -new -key key.pem -out server.csr -subj "/C=FI/ST=Turku/L=Turku/O=JaspnasCom/OU=IT Department/CN=reaktorbirdnest.jaspnas.com" && openssl x509 -req -days 100 -in server.csr -signkey key.pem -out cert.crt` to generate a self-signed certificate.

`docker compose up -d` in the root directory

## Description

The project consist of 4 docker containers running on a separate network using Docker compose. 

#### Backend

The backend module updates the database and serves information about drones to clients. It is written in golang.

#### Frontend

The frontend module is written in react and shows the data fetched from the backend.

#### Mongo

The project uses mongodb to store drones.

#### Yxorp

The project uses nginx as a reverse proxy to distribute traffic between the frontend and backend. 

### Tying it togehter

The frontend first requests data from the backend using a get request to the `/api/drones` endpoint. After this it connects to a websocket at `/api/websocket` that continuosly sends all drones caught violating the protected area. 
The valid (CA-signed) certificates are only stored on the reverse proxy server while `backend` and `frontend` use randomly generated certificates. 
