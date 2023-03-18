#!/bin/bash

number=${1:-$1}
EMBED=${2:-true}

# approve
if [[ "${EMBED}" == false ]]; then
    for i in $(seq 1 $number); do
        for j in $(seq 1 10); do 
            clustername=mc-$i-$j-external
            acceptcmd=$(clusteradm --kubeconfig hack/deploy/cert-multicluster-controlplane-external-$i/kubeconfig accept --clusters $clustername)
            echo $acceptcmd
        done 
    done 
else
    for i in $(seq 1 $number); do
        for j in $(seq 1 10); do 
            clustername=mc-$i-$j-embed 
            acceptcmd=$(clusteradm --kubeconfig hack/deploy/cert-multicluster-controlplane-embed-$i/kubeconfig accept --clusters $clustername)
            echo $acceptcmd
        done 
    done 
fi 
