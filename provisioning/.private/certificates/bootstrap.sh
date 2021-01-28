# !/bin/bash

# Create the main Certificate Authority certificate and key
./scripts/certificate-authority.sh \
-ca-key "ca-key.pem" \
-ca-cert "ca-cert.pem" \
-pass $CA_PASSWORD \
-common-name "Root-CA" \
-country "BR" \
-org-unit "Seratos" \
-org "Seratos" \
-locality "Belo-Horizonte" \
-state "Minas-Gerais" 

# Create the truststore for the kafka and zookeeper
./scripts/truststore.sh \
-ca-key "ca-key.pem" \
-ca-cert "ca-cert.pem" \
-out "truststore" \
-pass $TRUSTSTORE_PASSWORD \
-common-name "Root-CA" \
-country "BR" \
-org-unit "Seratos" \
-org "Seratos" \
-locality "Belo-Horizonte" \
-name "truststore"
