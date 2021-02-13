#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"


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

