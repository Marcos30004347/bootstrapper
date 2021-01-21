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

To start the system run:
    
    vagrant up

To suspend the system run:
    
    vagrant suspend

To reload the system run:
    
    vagrant reload

To ssh into one machine run:
    
    vagrant ssh $machine_name

To halt the system run:
    
    vagrant halt

To destroy the system run:
    
    vagrant destroy


TO DO:
1. deploy config