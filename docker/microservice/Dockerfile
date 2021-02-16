FROM seratos/alpine:latest

RUN apk --no-cache add ca-certificates wget unzip curl

RUN curl -o consul.zip https://releases.hashicorp.com/consul/1.4.4/consul_1.4.4_linux_amd64.zip
RUN unzip consul.zip
RUN mv consul /usr/bin/
RUN mkdir -p /etc/consul.d
RUN mkdir -p /var/consul
RUN curl -L https://getenvoy.io/cli | sudo bash -s -- -b /usr/local/bin
RUN getenvoy run standard:1.13.6 -- --version
RUN cp ~/.getenvoy/builds/standard/1.13.6/linux_glibc/bin/envoy /usr/local/bin/

RUN echo 'datacenter = "datacenter" \ndata_dir = "/var/consul" \n \n \nencrypt = "abcde12345" \n \n \nca_file = "$root/consul-ca.pem" \n \nverify_incoming = true \nverify_outgoing = true \nverify_server_hostname = true \n \nacl = { \n  enabled = true \n  default_policy = "deny" \n  enable_token_persistence = true \n} \n \nserver = false \n \nretry_join = ["190.120.88.11","190.120.88.12","190.120.88.13"] \nstart_join = ["190.120.88.11","190.120.88.12","190.120.88.13"] \n \nconnect { \n    enabled = true \n} \n' >> /etc/consul.d/config.hcl

COPY CONF_FILE ./etc/service/config.yml

RUN j2 /tmp/config.hcl.j2 /etc/service/config.yml >> /etc/consul.d/config.hcl
RUN j2 /tmp/consul.j2 /etc/service/config.yml >> /etc/init.d/consul

RUN rm /tmp/config.hcl.j2
RUN rm /tmp/consul.j2

ENTRYPOINT ENTRY