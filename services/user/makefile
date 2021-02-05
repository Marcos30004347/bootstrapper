all:
	if [ -f executable ]; then rm executable; fi
	go build -o executable main.go

run:
	./executable

container:
	cp -r ../../protos protos/
	docker build -t user-service .
	rm -rf protos/

compile-protos:
	if [ -d ./codegen ]; then rm -rf ./codegen; fi
	mkdir -p codegen
	for service in `ls -d ./protos/*`; do \
	  for serv_ver in `ls -d $$service/*`; do \
	    for proto in `ls -d $$serv_ver/*`; do \
		    protoc --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go_out=./codegen --go-grpc_out=./codegen $$proto; \
      done; \
    done; \
	done
	mv ./codegen/protos ./codegen/pb
