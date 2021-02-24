codegen:
	chmod 777 ./scripts/codegen.sh
	./scripts/codegen.sh

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build .


run:
	sudo env "PATH=${PATH}" go run . --etcd-servers localhost:2379 \
    --authentication-kubeconfig ${HOME}/.kube/config \
    --authorization-kubeconfig ${HOME}/.kube/config \
    --kubeconfig ${HOME}/.kube/config