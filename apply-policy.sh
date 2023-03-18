#!/bin/bash

for j in $(seq 1 3); do
     for i in $(seq 1 50); do 
        cat ./policy-example.yaml | sed "s/name: policy-limitrange/name: policy-limitrange-$i/g" | oc --kubeconfig /root/go/src/ycyaoxdu/multicluster-controlplane/hack/deploy/cert-multicluster-controlplane-external-$j/kubeconfig apply -f -
     done
done

for j in $(seq 1 1); do
   for i in $(seq 1 50); do 
      cat ./policy-example.yaml | sed "s/name: policy-limitrange/name: policy-limitrange-$i/g" | oc --kubeconfig /root/go/src/ycyaoxdu/multicluster-controlplane/hack/deploy/cert-multicluster-controlplane-embed-$j/kubeconfig apply -f -
   done
done
