#!/bin/bash

SOURCE="${BASH_SOURCE[0]}"

while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done

CURRENT_DIRECTORY="$( cd -P "$( dirname "$SOURCE" )" >/dev/null 2>&1 && pwd )"

# Create consul certificate authority
$CURRENT_DIRECTORY/../../scripts/certificate-authority.sh \
    -ca-key "consul-agent-ca-key.pem " \
    -ca-cert "consul-agent-ca.pem " \
    -pass $CA_PASSWORD \
    -common-name "Root-CA" \
    -country "BR" \
    -org-unit "Marcos" \
    -org "Marcos" \
    -locality "Belo-Horizonte" \
    -state "Minas-Gerais" 

# Create consul keyfile
openssl rand -base64 4 > keyfile.pem