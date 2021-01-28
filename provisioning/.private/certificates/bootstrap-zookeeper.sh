# !/bin/bash

# Create the keystore for zookeeper
./scripts/keystore.sh \
-truststore "truststore/truststore.jks" \
-out "zookeeper" \
-name "zookeeper-keystore-1" \
-ca-key "ca-key.pem" \
-ca-cert "ca-cert.pem" \
-pass $ZOOKEEPER_KEYSTORE_PASSWORD \
-common-name "elliot" \
-country "BR" \
-org-unit "Seratos" \
-org "Seratos" \
-state "MG" \
-locality "Belo-Horizonte" \
-dns "DNS:elliot,IP:192.168.50.10,DNS:localhost,IP:127.0.0.1"

