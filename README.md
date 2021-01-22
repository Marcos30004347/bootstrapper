## Workspace Template

Development enviroment with Kubernetes and Vagrant

This repo was designed to simulate a production environment, so system testing can be done more accurately.

Dependencies:
1. VirtualBox
2. Vangrant
3. Ansible

There is currently three VMs:
    k8s-master
    k8s-node-1
    k8s-node-2

### Vagrant:
To start the system run:
    
    vagrant up

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