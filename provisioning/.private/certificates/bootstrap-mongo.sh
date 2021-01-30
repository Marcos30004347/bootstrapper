# !/bin/bash

# Create the keyfile for mongodb
openssl rand -base64 756 > mongo/mongo-keyfile.key


# Create the Certificate Authority
openssl genrsa -out mongo/mongo-ca.key \
    -passout pass:$MONGO_CERTIFICATE_PASSWORD \
    -aes256 8192 
openssl req -passin pass:$MONGO_CERTIFICATE_PASSWORD -x509 -new -extensions v3_ca -key mongo/mongo-ca.key -days 3650 -out mongo/mongo-ca-pub.crt -config mongo.conf 


# Create a certificate for each replica member
openssl req -nodes -newkey rsa:4096 -sha256 \
    -keyout mongo/mongo1.key \
    -out mongo/mongo1.csr \
    -passout pass:$MONGO_CERTIFICATE_PASSWORD \
    -passin pass:$MONGO_CERTIFICATE_PASSWORD \
    -config mongo.conf

openssl x509 -req -passin pass:$MONGO_CERTIFICATE_PASSWORD -CA mongo/mongo-ca-pub.crt -CAkey mongo/mongo-ca.key \
  -in mongo/mongo1.csr -out mongo/mongo1.crt \
  -days 3650 -CAcreateserial -extensions SAN -extfile mongo.conf

cat mongo/mongo1.key mongo/mongo1.crt > mongo/mongod1.pem


# # Create a certificate for mongo client/drive4r
openssl req -nodes -passin pass:$MONGO_CERTIFICATE_PASSWORD \
    -passout pass:$MONGO_CERTIFICATE_PASSWORD -newkey rsa:4096 \
    -sha256 -keyout mongo/client.key \
    -out mongo/client.csr -config mongo.conf

openssl x509 -req -CA mongo/mongo-ca-pub.crt \
    -CAkey mongo/mongo-ca.key -in mongo/client.csr \
    -out mongo/client.crt -days 3650 \
    -passin pass:$MONGO_CERTIFICATE_PASSWORD \
    -CAcreateserial -extensions SAN -extfile mongo.conf

cat mongo/client.key mongo/client.crt > mongo/client.pem

