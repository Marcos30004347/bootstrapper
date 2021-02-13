#!/usr/bin/env bash

set -e

VALIDITY_IN_DAYS=3650

CA_CERT_FILE="ca-cert.pem"
CA_KEY_FILE="ca-key.pem"

KEYSTORE_FILENAME="keystore.jks"
TRUSTSTORE_FILENAME="truststore.jks"

TRUSTSTORE_WORKING_DIRECTORY=""
KEYSTORE_WORKING_DIRECTORY="keystore"
KEYSTORE_SIGN_REQUEST="cert-file"
KEYSTORE_SIGN_REQUEST_SRL="ca-cert.srl"
KEYSTORE_SIGNED_CERT="cert-signed"

CN=""
OU=""
O=""
L=""
S=""
C=""

while test $# -gt 0; do
      case "$1" in
        -common-name)
            shift
            CN=$1
            shift
            ;;
        -country)
            shift
            C=$1
            shift
            ;;
        -state)
            shift
            S=$1
            shift
            ;;
        -locality)
            shift
            L=$1
            shift
            ;;
        -org)
            shift
            O=$1
            shift
            ;;
        -org-unit)
            shift
            OU=$1
            shift
            ;;
        -out)
            shift
            TRUSTSTORE_WORKING_DIRECTORY=$1/$TRUSTSTORE_WORKING_DIRECTORY
            shift
            ;;
        -ca-cert)
            shift
            CA_CERT_FILE=$1
            shift
            ;;
        -ca-key)
            shift
            CA_KEY_FILE=$1
            shift
            ;;
        -pass)
            shift
            PASSWORD=$1
            shift
            ;;
        -name)
            shift
            TRUSTSTORE_FILENAME=$1".jks"
            shift
            ;;
        *)
            echo "$1 is not a recognized flag!"
            return 1;
            ;;
    esac
done  


if [ -z "$PASSWORD" ]; then
  echo "You should set a password with the -pass flag!";
  exit 1
fi


DNAME="CN=$CN, OU=$OU, O=$O, L=$L, S=$S, C=$C"
SUBJ_CA="/C=$C/ST=$S/L=$L/O=$O/OU=$OU/CN=$CN"

if [ -e "$TRUSTSTORE_WORKING_DIRECTORY/$TRUSTSTORE_FILENAME" ]; then
  rm $TRUSTSTORE_WORKING_DIRECTORY/$TRUSTSTORE_FILENAME
fi


trust_store_file=""

mkdir -p $TRUSTSTORE_WORKING_DIRECTORY

echo
echo " - $CA_KEY_FILE -- the private key used later to"
echo "   sign certificates"
echo " - $CA_CERT_FILE -- the certificate that will be"
echo "   stored in the trust store in a moment and serve as the certificate"
echo "   authority (CA). Once this certificate has been stored in the trust"
echo "   store, it will be deleted. It can be retrieved from the trust store via:"

echo
echo "Now the trust store will be generated from the certificate."
echo
keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $TRUSTSTORE_WORKING_DIRECTORY/$TRUSTSTORE_FILENAME \
-alias CARoot -import -file $CA_CERT_FILE

echo
echo "$TRUSTSTORE_WORKING_DIRECTORY/$TRUSTSTORE_FILENAME was created."

echo
echo "All done!"
echo