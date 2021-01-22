In this folder are all the playbooks for ansible.

Inside roles/ is all the playbook roles for installation of dependencies and configuration. The install-kubernetes playbook is set to not use dokcer as the container engine even though there also a install-docker role, kubernetes will beb configured to use contanerd.