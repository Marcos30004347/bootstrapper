#!/usr/bin/env bash

set -e

VALIDITY_IN_DAYS=3650

CA_CERT_FILE="ca-cert.pem"
CA_KEY_FILE="ca-key.pem"

KEYSTORE_FILENAME="keystore.jks"
TRUSTSTORE_FILENAME="truststore.jks"

KEYSTORE_WORKING_DIRECTORY=""
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
            KEYSTORE_WORKING_DIRECTORY=$1
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
        -truststore)
            shift
            TRUSTSTORE_FILENAME=$1
            shift
            ;;
        -name)
            shift
            NAME=$1
            KEYSTORE_FILENAME=$1.jks
            shift
            ;;
        -dns)
            shift
            DNS=$1
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


if [ -e "$KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME" ]; then
  rm $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME
fi

if [ -e "$KEYSTORE_SIGN_REQUEST" ]; then
  rm -rf $KEYSTORE_SIGN_REQUEST
fi

if [ -e "$KEYSTORE_SIGN_REQUEST_SRL" ]; then
  rm -rf $KEYSTORE_SIGN_REQUEST_SRL
fi

if [ -e "$KEYSTORE_SIGNED_CERT" ]; then
  rm -rf $KEYSTORE_SIGNED_CERT
fi

trust_store_file="$TRUSTSTORE_FILENAME"
trust_store_private_key_file="$CA_KEY_FILE"

echo
echo "Continuing with:"
echo " - trust store file:        $trust_store_file"
echo " - trust store private key: $trust_store_private_key_file"

mkdir -p $KEYSTORE_WORKING_DIRECTORY

echo
echo "Now, a keystore will be generated. Each broker and logical client needs its own"
echo "keystore. This script will create only one keystore. Run this script multiple"
echo "times for multiple keystores."
echo

# To learn more about CNs and FQDNs, read:
# https://docs.oracle.com/javase/7/docs/api/javax/net/ssl/X509ExtendedTrustManager.html

keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME \
  -alias localhost -validity $VALIDITY_IN_DAYS -genkey -keyalg RSA

echo
echo "'$KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME' now contains a key pair and a"
echo "self-signed certificate. Again, this keystore can only be used for one broker or"
echo "one logical client. Other brokers or clients need to generate their own keystores."

echo
echo "Fetching the certificate from the trust store and storing in $CA_CERT_FILE."
echo

keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $trust_store_file -export -alias CARoot -rfc -file $CA_CERT_FILE

echo
echo "Now a certificate signing request will be made to the keystore."
echo
keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias localhost \
  -certreq -file $KEYSTORE_SIGN_REQUEST

echo
echo "Now the trust store's private key (CA) will sign the keystore's certificate."
echo

openssl x509 -req -passin pass:$PASSWORD -CA $CA_CERT_FILE -CAkey $trust_store_private_key_file \
  -in $KEYSTORE_SIGN_REQUEST -out $KEYSTORE_SIGNED_CERT \
  -days $VALIDITY_IN_DAYS -CAcreateserial -extensions SAN -extfile <(printf "\n[SAN]\nsubjectAltName=$DNS")

echo
echo "Now the CA will be imported into the keystore."
echo
keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias CARoot \
  -import -file $CA_CERT_FILE

echo
echo "Now the keystore's signed certificate will be imported back into the keystore."
echo
keytool -noprompt -storepass $PASSWORD -dname "${DNAME}" -keystore $KEYSTORE_WORKING_DIRECTORY/$KEYSTORE_FILENAME -alias localhost -import \
  -file $KEYSTORE_SIGNED_CERT

echo
echo "All done!"
echo

rm $KEYSTORE_SIGN_REQUEST_SRL
rm $KEYSTORE_SIGN_REQUEST
rm $KEYSTORE_SIGNED_CERT
