NUMBER_OF_MACHINES = 3
ANSIBLE_RAW_SSH_ARGS = []

Vagrant.configure("2") do |config|
    # config.ssh.insert_key = false

    (1..NUMBER_OF_MACHINES-1).each do |machine_id|
        ANSIBLE_RAW_SSH_ARGS << "-o IdentityFile=#{Dir.pwd}/.vagrant/machines/machine#{machine_id}/virtualbox/private_key"
    end

    (1..NUMBER_OF_MACHINES).each do |i|
        # config.ssh.private_key_path = "/home/marcos/.vagrant.d/insecure_private_key"
        # config.ssh.forward_agent = true
        # config.ssh.username = 'username'
        config.vm.define "machine#{i}" do |machine|
            machine.vm.box = "ubuntu/xenial64"
            machine.vm.hostname  = "machine#{i}"
            machine.vm.network :private_network, ip: "190.120.88.1#{i}"
            machine.vm.provider "virtualbox" do |vb|
                vb.memory = "4096"
                vb.cpus = 2
            end

            # Boot all machines before provisioning
            if i == NUMBER_OF_MACHINES
                # machine.vm.provision :ansible do |ansible|
                #     ansible.limit           = "all"
                #     ansible.playbook        = "ansible/mongodb.yml"
                #     ansible.inventory_path  = "ansible/local-inventory.yml"
                #     ansible.raw_ssh_args = ANSIBLE_RAW_SSH_ARGS
                #     ansible.ask_become_pass = true
                # end

                machine.vm.provision :ansible do |ansible|
                    ansible.limit           = "all"
                    ansible.playbook        = "ansible/consul.yml"
                    ansible.inventory_path  = "ansible/local-inventory.yml"
                    ansible.raw_ssh_args = ANSIBLE_RAW_SSH_ARGS
                    ansible.ask_become_pass = true
                end
            end
        end


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
    #     ansible.playbook       = "ansible/zookeeper.yml"
    #     ansible.inventory_path = "ansible/local-inventory.yml"
    #     ansible.ask_become_pass = true
    # end

    # config.vm.provision "ansible" do |ansible|
    #     ansible.playbook       = "ansible/kafka.yml"
    #     ansible.inventory_path = "ansible/local-inventory.yml"
    #     ansible.ask_become_pass = true
    # end

    # config.vm.provision "ansible" do |ansible|
    #     ansible.playbook       = "ansible/kubernetes.yml"
    #     ansible.inventory_path = "ansible/local-inventory.yml"
    #     ansible.ask_become_pass = true
    # end


end