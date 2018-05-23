#!/bin/bash

function getcert {
	local name=$1 ; shift
        local authtypes=$*

	kubectl delete csr $name

	local newline=$'\n'
	usages=""
	for a in $authtypes
	do
		usages="$usages  - $a auth$newline"
	done

	cat <<EOF | kubectl create -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: $name
spec:
  groups:
  - system:authenticated
  request: $(cat $name.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
$usages
EOF

	kubectl certificate approve $name
	kubectl get csr $name -o jsonpath='{.status.certificate}' | base64 -D > $name.pem
}

function mksecret {
	local ns=$1
	local name=$2
	kubectl -n $ns delete secret $name
	kubectl -n $ns create secret tls $name --cert=$name.pem --key=$name-key.pem
}

function genkey {
	local name=$1 ; shift
	local cn=$1 ; shift
	local hosts=$*

	if [ -z "$hosts" ]; then
	# Generate a new private key to represent this service, and a CSR for the public key.
		json=$(cfssl genkey - <<EOF
{ "CN": "$cn", "key": { "algo": "rsa", "size": 2048 } }
EOF
)
	else
		local hh=$(echo $hosts | sed -e 's/ /","/g')
		json=$(cfssl genkey - <<EOF
{ "hosts": [ "$hh" ], "CN": "$cn", "key": { "algo": "rsa", "size": 2048 } }
EOF
)
	fi

	echo $json | jq -r '.key' > $name-key.pem
	echo $json | jq -r '.csr' > $name.csr
}


#genkey grpc-whoami-client grpc-whoami-client
#getcert grpc-whoami-client client
#mksecret default grpc-whoami-client

genkey grpc-whoami grpcwhoami-service grpcwhoami.default.svc.cluster.local grpcwhoami.default.svc grpcwhoami.default grpcwhoami
getcert grpc-whoami server client
mksecret default grpc-whoami
