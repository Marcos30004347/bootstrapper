codegen:
	chmod 777 ./scripts/codegen.sh
	./scripts/codegen.sh

build:
	docker build -t marcos30004347/kubeapi .

run:
	sudo env "PATH=${PATH}" go run . --etcd-servers localhost:2379 \
    --authentication-kubeconfig ${HOME}/.kube/config \
    --authorization-kubeconfig ${HOME}/.kube/config \
    --kubeconfig ${HOME}/.kube/config


deploy: build
	# eval $(minikube docker-env)
	kubectl apply -f ./k8s/ns.yaml
	kubectl apply -f ./k8s/
