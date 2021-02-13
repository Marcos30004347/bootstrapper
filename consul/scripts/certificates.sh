set -e

# Cert subj parameters
COUNTRY="BR"
STATE="MinasGerais"
ORG="Marcos"
CN="RootCA"
EMAIL="teste@example.org"
SUBJ_BASE="/C=${COUNTRY}/ST=${STATE}/O=${ORG}/emailAddress=${EMAIL}"

# Consul parameters
CONSUL_DC="datacenter"
CONSUL_DOMAIN="consul"

# Create the -subj arg for each cert
CA_SUBJ="/CN=${CN}${SUBJ_BASE}"
SERVER_SUBJ="/CN=server.${CONSUL_DC}.${CONSUL_DOMAIN}${SUBJ_BASE}"
CLIENT_SUBJ="/CN=client.${CONSUL_DC}.${CONSUL_DOMAIN}${SUBJ_BASE}"

# Key specifications
KEY_SIZE=4096

CONSUL_SERVERS_COUNT=3
CONSUL_CLIENTS_COUNT=3

# Get current directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
CURRENT_DIRECTORY="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"


if [ -d $CURRENT_DIRECTORY/../certificates ]; then find $CURRENT_DIRECTORY/../certificates/ -name "*.pem" -type f -delete; fi
if [ -d $CURRENT_DIRECTORY/../certificates ]; then find $CURRENT_DIRECTORY/../certificates/ -name "*.old" -type f -delete; fi
if [ -d $CURRENT_DIRECTORY/../certificates ]; then rm -rf $CURRENT_DIRECTORY/../certificates/state; fi

mkdir -p $CURRENT_DIRECTORY/../certificates
CERT_OUT_DIR=$CURRENT_DIRECTORY/../certificates


# Create the directories/files needed for the CA
# mkdir -p files
mkdir -p $CERT_OUT_DIR/state
echo "000a" > $CERT_OUT_DIR/serial.pem
cd ../certificates
touch state/certindex

# Generate a CA key and certificate
openssl genrsa -out $CERT_OUT_DIR/consul-ca-key.pem "$KEY_SIZE"
openssl req -x509 -new -nodes -key $CERT_OUT_DIR/consul-ca-key.pem -subj "$CA_SUBJ" -days 3650 -out $CERT_OUT_DIR/consul-ca.pem -sha256

for ((i = 0; i < $CONSUL_SERVERS_COUNT; i++)); do
    # Generate keys and certificates for Consul server agents
    openssl genrsa -out $CERT_OUT_DIR/$CONSUL_DC-server-$i-key.pem "$KEY_SIZE";
    openssl req -new -key $CERT_OUT_DIR/$CONSUL_DC-server-$i-key.pem -subj "$SERVER_SUBJ" -out $CERT_OUT_DIR/$CONSUL_DC-server-$i.pem -sha256;
    openssl ca -batch -config $CERT_OUT_DIR/ca.conf -notext -in $CERT_OUT_DIR/$CONSUL_DC-server-$i.pem -out $CERT_OUT_DIR/$CONSUL_DC-server-$i.pem;
done

for ((i = 0; i < $CONSUL_CLIENTS_COUNT; i++)); do
    # Generate keys and certificates for Consul client agents
    openssl genrsa -out $CERT_OUT_DIR/$CONSUL_DC-client-$i-key.pem "$KEY_SIZE";
    openssl req -new -key $CERT_OUT_DIR/$CONSUL_DC-client-$i-key.pem -subj "$CLIENT_SUBJ" -out $CERT_OUT_DIR/$CONSUL_DC-client-$i.pem -sha256;
    openssl ca -batch -config $CERT_OUT_DIR/ca.conf -notext -in $CERT_OUT_DIR/$CONSUL_DC-client-$i.pem -out $CERT_OUT_DIR/$CONSUL_DC-client-$i.pem;
done

# Create gossip consul keyfile
openssl rand -base64 32 > $CERT_OUT_DIR/gossip-key.pem

cd ..


rm -rf $CERT_OUT_DIR/state
rm -rf $CERT_OUT_DIR/serial.pem
rm -rf $CERT_OUT_DIR/serial.pem.old
