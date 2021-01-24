## Application core

In this folder are all services implementation that are the core of the application. 

The deploy.yaml file is the kubernetes deployment definition that tells k8s how to deploy the application inside the cluster.

To deploy, from the root folder, run:

    kubectl apply -f core/deploy.yaml 