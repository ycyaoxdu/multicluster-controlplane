#!/bin/bash

number=${1:-$1}
EMBED=${2:-true}

export KUBECTL=oc

if [[ "${EMBED}" == false ]]; then
    # external etcd
    ## 1. deploy etcd
    export ETCD_NS=multicluster-etcd
    oc delete ns $ETCD_NS

    ## 2. controlplane
    for i in $(seq 1 "${number}"); do
        namespace=multicluster-controlplane-external-$i
        oc delete ns $namespace
    done
else 
    # embedetcd
    for i in $(seq 1 "${number}"); do
        namespace=multicluster-controlplane-embed-$i
        oc delete ns $namespace
    done
fi

