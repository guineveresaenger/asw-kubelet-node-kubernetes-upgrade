*This project aids in updating the Kubernetes version of AWS kubelet nodes with k2cli config.*

To update forward/backward the Kubernetes version of an existing cluster, change the Kubernetes version in k2cli config file under deployment.kubeConfig.version and run the cluster up to overwrite the cluster version.

Run the bash script in this project to drain and delete non-master nodes from Kubernetes and AWS. When AWS auto-scaling group brings back the nodes, they will have the updated Kubernetes version.
