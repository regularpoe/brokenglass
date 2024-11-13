# wiretrap

openssl genrsa -out server.key 2048

openssl req -new -key server.key -out server.csr

openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

openssl pkcs12 -export -out identity.pfx -inkey key.pem -in cert.pem

openssl s_client -connect 127.0.0.1:2408
