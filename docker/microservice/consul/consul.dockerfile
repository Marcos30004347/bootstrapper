RUN curl -o consul.zip https://releases.hashicorp.com/consul/1.4.4/consul_1.4.4_linux_amd64.zip
RUN unzip consul.zip
RUN mv consul /usr/bin/
RUN mkdir -p /etc/consul.d
RUN mkdir -p /var/consul
RUN curl -L https://getenvoy.io/cli | sudo bash -s -- -b /usr/local/bin
RUN getenvoy run standard:1.13.6 -- --version
RUN cp ~/.getenvoy/builds/standard/1.13.6/linux_glibc/bin/envoy /usr/local/bin/

RUN eval $(echo '\
datacenter = "{{ discovery.datacenter }}" \n \
data_dir = "/var/consul" \n \
\n \
{% if discovery.encrypt is defined %} \n \
encrypt = "{{ discovery.encrypt }}" \n \
{% endif %} \n \
ca_file = "{{ discovery.ca }}" \n \
\n \
verify_incoming = true \n \
verify_outgoing = true \n \
verify_server_hostname = true \n \
\n \
acl = { \n \
  enabled = true \n \
  default_policy = "deny" \n \
  enable_token_persistence = true \n \
} \n \
\n \
server = false \n \
\n \
retry_join = [{% for host in discovery.connect %}"{{ host }}"{% if not loop.last %},{% endif %}{% endfor %}] \n \
start_join = [{% for host in discovery.connect %}"{{ host }}"{% if not loop.last %},{% endif %}{% endfor %}] \n \
\n \
connect { \n \
    enabled = true \n \
}\n\
')
