Vagrant.configure("2") do |config|
    config.ssh.insert_key = false

    # Provider set to virtual box, each machine will
    # have 2GB of memory and a CPU with 2 virtual cores
    config.vm.box = "ubuntu/xenial64"

    config.vm.provider "virtualbox" do |virtualbox|
        virtualbox.memory = 4096 # 2GB
        virtualbox.cpus = 2 # 2vCPU
    end

    config.vm.define "mongo1" do |mongo1|
        mongo1.vm.hostname  = "mongo1"
        mongo1.vm.network :private_network, ip: "190.120.88.11"
    end

    config.vm.define "mongo2" do |mongo2|
        mongo2.vm.hostname  = "mongo2"
        mongo2.vm.network :private_network, ip: "190.120.88.12"
    end

    config.vm.define "mongo3" do |mongo3|
        mongo3.vm.hostname  = "mongo3"
        mongo3.vm.network :private_network, ip: "190.120.88.13"
    end

    # # Elliot machine
    # config.vm.define "elliot" do |master|
    #     master.vm.hostname  = "elliot"
    #     master.vm.network "private_network", ip: "192.168.50.10"
    #     master.vm.disk :disk, size: "2GB", primary: true
    # end

    # # Darlene machine
    # config.vm.define "darlene" do |master|
    #     master.vm.hostname  = "darlene"
    #     master.vm.network "private_network", ip: "170.154.32.13"
    #     master.vm.disk :disk, size: "2GB", primary: true
    # end

    # # Tyrell machine
    # config.vm.define "tyrell" do |master|
    #     master.vm.hostname  = "tyrell"
    #     master.vm.network "private_network", ip: "140.320.42.15"
    #     master.vm.disk :disk, size: "2GB", primary: true
    # end

    # Angela machine
    # config.vm.define "angela" do |angela|
    #     angela.vm.hostname  = "angela"
    #     angela.vm.network "private_network", ip: "165.120.88.2"
    #     angela.vm.disk :disk, size: "10GB", primary: true
    # end

    # config.vm.provision "ansible" do |ansible|
    #     ansible.playbook       = "provisioning/zookeeper.yml"
    #     ansible.inventory_path = "provisioning/inventory.yml"
    #     ansible.ask_become_pass = true
    # end

    # config.vm.provision "ansible" do |ansible|
    #     ansible.playbook       = "provisioning/kafka.yml"
    #     ansible.inventory_path = "provisioning/inventory.yml"
    #     ansible.ask_become_pass = true
    # end

    # config.vm.provision "ansible" do |ansible|
    #     ansible.playbook       = "provisioning/kubernetes.yml"
    #     ansible.inventory_path = "provisioning/inventory.yml"
    #     ansible.ask_become_pass = true
    # end

    config.vm.provision "ansible" do |ansible|
        ansible.extra_vars      = { ansible_ssh_user: 'vagrant' }
        ansible.raw_arguments  = [
            "--private-key=/home/marcos/.vagrant.d/insecure_private_key"
        ]
        ansible.playbook        = "provisioning/mongodb.yml"
        ansible.inventory_path  = "provisioning/inventory.yml"
        ansible.ask_become_pass = true
    end

end