#!/bin/bash

mkdir -p ssl

cat << EOF > ssl/req.cnf
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[v3_req]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names
[alt_names]
# 此处填写需要在证书中添加的域名白名单，格式为 DNS.1、DNS.2、DNS.3，依次类推
DNS.1 = ldap.example.org
# 此处填写需要在证书中添加的IP白名单，格式为 IP.1、IP.2、IP.3，依次类推
IP.1 = 192.168.1.179
IP.2 = 172.17.0.2
EOF

openssl genrsa -out ssl/ca-key.pem 2048
openssl req -x509 -new -nodes -key ssl/ca-key.pem -days 8000 -out ssl/ca.pem -subj "/CN=kube-ca"

openssl genrsa -out ssl/key.pem 2048
openssl req -new -key ssl/key.pem -out ssl/csr.pem -subj "/CN=kube-ca" -config ssl/req.cnf
openssl x509 -req -in ssl/csr.pem -CA ssl/ca.pem -CAkey ssl/ca-key.pem -CAcreateserial -out ssl/cert.pem -days 8000 -extensions v3_req -extfile ssl/req.cnf
