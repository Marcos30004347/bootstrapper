NUMBER_OF_MACHINES = 3
ANSIBLE_RAW_SSH_ARGS = []

Vagrant.configure("2") do |config|
    # config.ssh.insert_key = false

    (1..NUMBER_OF_MACHINES-1).each do |machine_id|
        ANSIBLE_RAW_SSH_ARGS << "-o IdentityFile=#{Dir.pwd}/.vagrant/machines/machine#{machine_id}/virtualbox/private_key"
    end

    (1..NUMBER_OF_MACHINES).each do |i|
        config.vm.define "machine#{i}" do |machine|
            machine.vm.box = "ubuntu/xenial64"
            machine.vm.hostname  = "machine#{i}"
            machine.vm.network :private_network, ip: "190.120.88.1#{i}"
            machine.vm.provider "virtualbox" do |vb|
                vb.memory = "4096"
                vb.cpus = 2
            end

            if i == NUMBER_OF_MACHINES
                # machine.vm.provision :ansible do |ansible|
                #     ansible.limit           = "all"
                #     ansible.playbook        = "kubernetes/ansible/playbook.yml"
                #     ansible.inventory_path  = "kubernetes/ansible/inventory.yml"
                #     ansible.raw_ssh_args = ANSIBLE_RAW_SSH_ARGS
                #     ansible.ask_become_pass = true
                # end
                machine.vm.provision :ansible do |ansible|
                    ansible.limit           = "all"
                    ansible.playbook        = "consul/ansible/playbook.yml"
                    ansible.inventory_path  = "consul/ansible/inventory.yml"
                    ansible.raw_ssh_args = ANSIBLE_RAW_SSH_ARGS
                    ansible.ask_become_pass = true
                end
            end
        end


    end
end