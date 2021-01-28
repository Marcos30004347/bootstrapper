Vagrant.configure("2") do |config|
    config.ssh.insert_key = false

    # Provider set to virtual box, each machine will
    # have 2GB of memory and a CPU with 2 virtual cores
    config.vm.box = "ubuntu/xenial64"

    config.vm.provider "virtualbox" do |virtualbox|
        virtualbox.memory = 4096 # 2GB
        virtualbox.cpus = 2 # 2vCPU
    end

    # Elliot machine
    config.vm.define "elliot" do |master|
        master.vm.hostname  = "elliot"
        master.vm.network "private_network", ip: "192.168.50.10"
    end

    # Darlene machine
    config.vm.define "darlene" do |master|
        master.vm.hostname  = "darlene"
        master.vm.network "private_network", ip: "170.154.32.13"
    end

    # Tyrell machine
    config.vm.define "tyrell" do |master|
        master.vm.hostname  = "tyrell"
        master.vm.network "private_network", ip: "140.320.42.15"
    end

    config.vm.provision "ansible" do |ansible|
        ansible.playbook       = "provisioning/zookeeper.yml"
        ansible.inventory_path = "provisioning/inventory.yml"
        ansible.ask_become_pass = true
    end

    config.vm.provision "ansible" do |ansible|
        ansible.playbook       = "provisioning/kafka.yml"
        ansible.inventory_path = "provisioning/inventory.yml"
        ansible.ask_become_pass = true
    end

    config.vm.provision "ansible" do |ansible|
        ansible.playbook       = "provisioning/kubernetes.yml"
        ansible.inventory_path = "provisioning/inventory.yml"
        ansible.ask_become_pass = true
    end
end