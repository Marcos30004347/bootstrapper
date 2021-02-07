zookeeper_url can be equal to localhost:2181 if zookeeper is running locally.

replica_factor is a number between 1 and the cout off kafka nodes running.

To list the kafka topics inside the machine run:

    /opt/kafka/bin/kafka-topics.sh --list --zookeeper $zookeeper_url

To list the kafka topics inside the machine run:

    /opt/kafka/bin/kafka-topics.sh --create --zookeeper $zookeeper_url --replication-factor $replica_factor --partitions 20 --topic $topic

To describe a topic:

    bin/kafka-topics.sh --zookeeper $zookeeper_url --topic $topic --describe

To remove a kafka topic, delete.topic.enable=true
 shoud be enabled in the server.properties and you should run:

    bin/kafka-topics.sh --zookeeper $zookeeper_url --topic $topic --delete


TODO:
1. https://developer.ibm.com/components/kafka/tutorials/kafka-authn-authz/#:~:text=Kafka%20manages%20and%20enforces%20authorization,that%20stores%20ACLs%20in%20ZooKeeper.
2. https://docs.confluent.io/platform/current/kafka/authorization.html
3. http://kafka.apache.org/090/documentation.html#zk_authz
4. https://docs.confluent.io/platform/current/kafka/authentication_ssl.html