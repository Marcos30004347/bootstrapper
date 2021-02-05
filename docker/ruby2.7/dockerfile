FROM ubuntu

# WORKDIR /task-service
# COPY . /task-service

RUN apt-get update
RUN apt-get install -y gcc
RUN apt-get install -y openssl
RUN apt-get install -y make
RUN apt-get install -y ruby2.7
RUN apt-get install -y ruby-dev

RUN gem install bundler
RUN gem install json
RUN gem install mongo
RUN gem install grpc
RUN gem install grpc-tools

# RUN bundle install
# RUN make setup
# RUN ls
# EXPOSE 50051
# ENTRYPOINT [ "make", "run" ]
