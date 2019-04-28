#!/bin/bash

# You need to change these constants based on your GOPATH
CERTS="/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/"
CLIENT_CERT_NAME="client1.crt"
CLIENT_KEY_NAME="client1.key"
CA_CERT_NAME="onfca.crt"

./gnishell    -gni_address "localhost:50051" \
       -target_address "localhost:10161" \
       -timeout "5" \
       -proto   "" \
       -rpc "fetch" \
       -rpc_type "gnmi:capability" \
       -target_name "Test-onos-config" \
       -client_crt $CERTS$CLIENT_CERT_NAME \
       -client_key $CERTS$CLIENT_KEY_NAME \
       -ca_crt $CERT$CA_CERT_NAME \
       
./gnishell    -gni_address "localhost:50051" \
       -target_address "localhost:10161" \
       -timeout "5" \
       -proto   "path: <elem: <name: 'system'> elem:<name:'config'> elem: <name: 'hostname'>>" \
       -rpc "fetch" \
       -rpc_type "gnmi:get" \
       -target_name "Test-onos-config" \
       -client_crt $CERTS$CLIENT_CERT_NAME \
       -client_key $CERTS$CLIENT_KEY_NAME \
       -ca_crt $CERT$CA_CERT_NAME \