#!/bin/bash

KEY_FILE=key.pem
CERT_FILE=cert.pem

# Generate the keypair
openssl genrsa -out $KEY_FILE 2048
openssl req -key $KEY_FILE -new -x509 -days 3650 -out $CERT_FILE
