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
    config.vm.define "elliot" do |master|
        master.vm.define "elliot"
        master.vm.box = "ubuntu/xenial64"
        # can be acesseb both by the ip than by the http://elliot
        master.vm.network "private_network", ip: "192.168.50.10"
        master.vm.hostname = "elliot"

        master.vm.provision "ansible" do |ansible|
            ansible.groups = {
                "k8smasters" => ["elliot"],
                "zookeepernodes" => ["elliot"],
                "kafkanodes" => ["elliot"],
            }
            ansible.playbook = "provisioning/elliot.yml"
            ansible.ask_become_pass = true

            ansible.extra_vars = {
                zookeeper_id: 1,
                node_ip: "192.168.50.10",
                home_dir: "/home/vagrant",
                linux_distro: "xenial",
                pod_network_cidr: "192.168.0.0/16",
            }

            ansible.host_vars = {
                "elliot" => {
                    zookeeper_id: 1,
                    kafka_broker_id: 1,
                    private_ip: "192.168.50.10",
                },
            }

        end
    end

    # zookeeper_id

    #########################################
    ###### Define the Kubernetes Nodes ######
    #########################################
    # (1..KUBERNETES_NODES_COUNT).each do |i|
    #     config.vm.define "k8snode#{i}" do |node|
    #         node.vm.define "k8snode#{i}"
    #         node.vm.box = "ubuntu/xenial64"

    #         # can be acesseb both by the ip than by the http://k8node#{i}
    #         node.vm.network "private_network", ip: "192.168.50.#{i + 10}"
    #         node.vm.hostname = "k8snode#{i}"
            
    #         node.vm.provision "ansible" do |ansible|
    #             ansible.groups = {
    #                 "k8snodes" => ["k8snode#{i}"],
    #             }
    #             ansible.playbook = "provisioning/node.playbook.yml"
    #             ansible.ask_become_pass = true
    #             ansible.extra_vars = {
    #                 linux_distro: "xenial",
    #                 node_ip: "192.168.50.#{i + 10}",
    #                 home_dir: "/home/vagrant",
    #             }
    #         end
    #     end
    # end
end