#!/bin/bash

number=${1:-$1}
EMBED=${2:-true}

export KUBECTL=oc

if [[ "${EMBED}" == false ]]; then
    # external etcd
    ## 1. deploy etcd
    export ETCD_NS=multicluster-etcd
    oc create ns $ETCD_NS
    make deploy-etcd 
    ## 2. controlplane
    for i in $(seq 1 "${number}"); do
        namespace=multicluster-controlplane-external-$i
        oc create ns $namespace
        export HUB_NAME=$namespace
        make deploy-with-external-etcd
    done
else 
    # embedetcd
    for i in $(seq 1 "${number}"); do
        namespace=multicluster-controlplane-embed-$i
        oc create ns $namespace
        export HUB_NAME=$namespace
        make deploy
    done
fi
