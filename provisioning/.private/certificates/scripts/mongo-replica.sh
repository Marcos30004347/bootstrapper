
OUTPUT=""
DNS=""
VALIDITY_IN_DAYS=3650

while test $# -gt 0; do
      case "$1" in
        -out)
            shift
            OUTPUT=$1
            shift
            ;;
        -name)
            shift
            NAME=$1
            shift
            ;;
        -pass)
            shift
            PASSWORD=$1
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
        -conf)
            shift
            CONFIG=$1
            shift
            ;;
        *)
            echo "$1 is not a recognized flag!"
            return 1;
            ;;
    esac
done  


openssl req -nodes -passin pass:$PASSWORD -passout pass:$PASSWORD -newkey rsa:4096 -sha256 -keyout $OUTPUT/$NAME.key -out $OUTPUT/$NAME.csr -config $CONFIG 

openssl x509 -req -passin pass:$PASSWORD -CA $CA_CERT_FILE -CAkey $CA_KEY_FILE \
  -in $OUTPUT/$NAME.csr -out $OUTPUT/$NAME.crt \
  -days $VALIDITY_IN_DAYS -CAcreateserial -extensions SAN -extfile $CONFIG

