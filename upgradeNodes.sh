#!/bin/bash

kubectl get no | awk '!/Sched|NAM/{print $1}' | while read line; do kubectl drain $line --force --ignore-daemonsets; kubectl delete node $line; var=$(go run formatNode.go --nodename=$line); aws ec2 terminate-instances --instance-ids $var; done
