## Workspace Template
TODO:
1. istio
2. sql
3. mongo
4. refactoring
5. security
   

Development enviroment with Kubernetes and Vagrant

This repo was designed to simulate a production environment, so system testing can be done more accurately.

Dependencies:
1. VirtualBox
2. Vangrant
3. Ansible
4. openssl
There is currently three VMs:
    k8s-master
    k8s-node-1
    k8s-node-2

### Setup

You need to export some variables before deploying the enviroment, we recomend you to install direnv(https://direnv.net/) and setup a .envrc file exporting the following variables:

1. CA_PASSWORD                      - the password used to create certificate authority
1. TRUSTSTORE_PASSWORD              - the password used to create the truststore certificate
2. KAFKA_USER_PASSWORD              - hashed password for the kafkaadmin user
3. KAFKA_KEYSTORE_PASSWORD          - the password used to create the keystore certificate
4. ZOOKEEPER_USER_PASSWORD          - hashed password for the zookeeperadmin user
5. ZOOKEEPER_KEYSTORE_PASSWORD      - the password used to create the keystore certificate

The .envrc file is in the .gitignore.

### Vagrant:
To start the system run:
    
    vagrant up --no-provision

To setup all the system machines run:
    
    vagrant provision

To suspend the system run:
    
    vagrant suspend

To reload the system run:
    
    vagrant reload

To ssh into one machine run:
    
    vagrant ssh <"machine_name">

To halt the system run:
    
    vagrant halt

To destroy the system run:
    
    vagrant destroy


### Kubernetes:
To list the nodes in the system
run:
    
    $ kubectl get nodes

To describe some entity inside the cluster run:

    $ kubectl describe <node | pod | deployment> <"name">

To get the pods run:
    
    $ kubectl get pods -o wide

You can also get the kubernetes exclusive pods with:

    $ kubectl get pods -n kube-system

To run a image inside the cluster run:

    $ kubectl run <"deployment_name"> --image <"image_name">

To scale something run:

    $ kubectl scale deployment <"deployment_name"> --replicas=<"count">

To remove a deployment run:

    $ kubectl delete deployment <"deployment_name">