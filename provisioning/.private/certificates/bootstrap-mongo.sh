# !/bin/bash

# Create the keyfile for mongodb
openssl rand -base64 756 > mongo/mongo-keyfile

# Create the Certificate Authority
openssl genrsa -out mongo-ca.key -passout pass:$MONGO_CERTIFICATE_PASSWORD -aes256 8192 
openssl req -passin pass:$MONGO_CERTIFICATE_PASSWORD -x509 -new -extensions v3_ca -key mongo-ca.key -days 3650 -out mongo-ca-pub.crt -config mongo.conf 

# Create a certificate for each replica member
openssl req -nodes \ 
    -newkey rsa:4096 -sha256 \
    -keyout mongo/replicas/mongo1.key \
    -out mongo/replicas/mongo1.csr \
    -passout pass:$MONGO_CERTIFICATE_PASSWORD \
    -passin pass:$MONGO_CERTIFICATE_PASSWORD \
    -config mongo.conf

# openssl x509 -req -passin pass:$MONGO_CERTIFICATE_PASSWORD -CA mongo-ca-pub.crt -CAkey mongo-ca.key \
#   -in mongo/replicas/mongo1.csr -out mongo/replicas/mongo1.crt \
#   -days 3650 -CAcreateserial -extensions SAN -extfile mongo.conf

# # Create a certificate for mongo client/drive4r
# echo "asdasdasdasdasdasdASDDSAd"
# echo "asdasdasdasdasdasdASDDSAd"
# echo "asdasdasdasdasdasdASDDSAd"
# echo "asdasdasdasdasdasdASDDSAd"
# openssl req -nodes -passin pass:$PASSWORD \
#     -passout pass:$PASSWORD -newkey rsa:4096 \
#     -sha256 -keyout mongo/client/client.key \
#     -out mongo/client/client.csr -config mongo.conf

# openssl x509 -req -passin pass:$PASSWORD -CA mongo-ca-pub.crt \
#     -CAkey mongo-ca.key -in mongo/client/client.csr \
#     -out mongo/client/client.crt -days 3650 \
#     -CAcreateserial -extensions SAN -extfile mongo.conf

# cat mongo/client/client.key mongo/client/client.crt > mongo/client/client.pem


# # Create the certificate for one mongo instance
# ./scripts/mongo-replica.sh \
# -ca-cert "mongo-ca-pub.crt" \
# -ca-key "mongo-ca.key" \
# -out "mongo/replicas" \
# -name "mongo1" \
# -common-name "elliot" \
# -pass $MONGO_CERTIFICATE_PASSWORD \
# -country "BR" \
# -state "MG" \
# -org-unit "Seratos" \
# -org "Seratos" \
# -locality "Belo-Horizonte" \
# -conf "mongo.conf" \
# -dns "DNS:elliot,DNS:angela,IP:165.120.88.2,IP:192.168.50.10,DNS:localhost,IP:127.0.0.1"

# ./scripts/mongo-client.sh \
# -ca-cert "mongo-ca-pub.crt" \
# -ca-key "mongo-ca.key" \
# -out "mongo/client" \
# -name "mongo1" \
# -common-name "elliot" \
# -pass $MONGO_CERTIFICATE_PASSWORD \
# -country "BR" \
# -state "MG" \
# -org-unit "Seratos" \
# -org "Seratos" \
# -locality "Belo-Horizonte" \
# -conf "mongo.conf" \
# -dns "DNS:elliot,DNS:angela,IP:165.120.88.2,IP:192.168.50.10,DNS:localhost,IP:127.0.0.1"
