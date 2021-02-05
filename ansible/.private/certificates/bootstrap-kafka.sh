# !/bin/bash

# Create the keystore for kafka
./scripts/keystore.sh \
-truststore "truststore/truststore.jks" \
-out "kafka" \
-name "kafka-keystore-1" \
-ca-key "ca-key.pem" \
-ca-cert "ca-cert.pem" \
-pass $KAFKA_KEYSTORE_PASSWORD \
-common-name "elliot" \
-country "BR" \
-state "MG" \
-org-unit "Seratos" \
-org "Seratos" \
-locality "Belo-Horizonte" \
-dns "DNS:elliot,IP:192.168.50.10,DNS:localhost,IP:127.0.0.1"
