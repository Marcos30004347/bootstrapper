#####################################
#### Vangrant configuration file ####
#####################################

# The vagrant file version number.
VAGRANT_FILE_VERSION="2"

# The count of nodes that should be instanciated with the master node.
KUBERNETES_NODES_COUNT=2


Vagrant.configure(VAGRANT_FILE_VERSION) do |config|
    config.ssh.insert_key = false

    # Provider set to virtual box, each machine will
    # have 2GB of memory and a CPU with 2 virtual cores
    config.vm.provider "virtualbox" do |virtualbox|
        virtualbox.memory = 4096 # 2GB
        virtualbox.cpus = 2 # 2vCPU
    end


    #########################################
    ### Define the Kubernetes Master Node ###
    #########################################
    config.vm.define "k8s-master" do |master|
        master.vm.box = "ubuntu/xenial64"

        # can be acesseb both by the ip than by the http://k8s-master
        master.vm.network "private_network", ip: "192.168.50.10"
        master.vm.hostname = "k8s-master"

        master.vm.provision "ansible", run: 'always' do |ansible|
            ansible.playbook = "ansible/k8s-master.playbook.yml"
            ansible.ask_become_pass = true
            ansible.extra_vars = {
                node_ip: "192.168.50.10",
                linux_distro: "xenial",
                home_dir: "/home/vagrant",
                pod_network_cidr: "192.168.0.0/16"
            }
        end
    end

    
    #########################################
    ###### Define the Kubernetes Nodes ######
    #########################################
    # (1..KUBERNETES_NODES_COUNT).each do |i|
    #     config.vm.define "k8s-node-#{i}" do |node|
    #         node.vm.box = "ubuntu/xenial64"
    
    #         # can be acesseb both by the ip than by the http://k8-node-#{i}
    #         node.vm.network "private_network", ip: "192.168.50.#{i + 10}"
    #         node.vm.hostname = "k8s-node-#{i}"
            
    #         node.vm.provision "ansible", run: 'always' do |ansible|
    #             ansible.playbook = "ansible/k8s-node.playbook.yml"
    #             ansible.ask_become_pass = true
    #             ansible.extra_vars = {
    #                 node_ip: "192.168.50.#{i + 10}",
    #                 linux_distro: "xenial",
    #                 home_dir: "/home/vagrant",
    #             }
    #         end
    #     end
    # end
end