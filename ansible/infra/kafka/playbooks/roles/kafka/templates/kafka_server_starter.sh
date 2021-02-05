#!/bin/sh
echo -e "Starting 9091 Server"
nohup sh /opt/kafka/bin/kafka-server-start.sh /opt/kafka/config/server.properties &