#!/usr/bin/env bash

set -e

CA_CERT_FILE="ca-cert.pem"
CA_KEY_FILE="ca-key.pem"
VALIDITY_IN_DAYS=3650

CN=""
OU=""
O=""
L=""
S=""
C=""

PASSWORD=123456

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
        *)
            echo "$1 is not a recognized flag!"
            return 1;
            ;;
    esac
done 

SUBJ_CA="/C=$C/ST=$S/L=$L/O=$O/OU=$OU/CN=$CN"


if [ -e "$CA_CERT_FILE" ]; then
  rm -rf $CA_CERT_FILE
fi

if [ -e "$CA_KEY_FILE" ]; then
  rm -rf $CA_KEY_FILE
fi



echo
echo "OK, we'll generate a trust store and associated private key."
echo
echo "First, the private key."

openssl req -new -x509 -keyout $CA_KEY_FILE -out $CA_CERT_FILE -subj $SUBJ_CA -passout pass:$PASSWORD -passin pass:$PASSWORD -days $VALIDITY_IN_DAYS

trust_store_private_key_file="$CA_KEY_FILE"

echo
echo "All done!"
echo

