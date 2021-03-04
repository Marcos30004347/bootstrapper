deps:
	go get -t .

startup:
	chmod 777 ./scripts/minikube-startup.sh
	./scripts/minikube-startup.sh

shutdown:
	chmod 777 ./scripts/minikube-shutdown.sh
	./scripts/minikube-shutdown.sh

build:
	docker build -t marcos30004347/kubeapi .

codegen:
	chmod 777 ./scripts/codegen.sh
	./scripts/codegen.sh

deploy:
	kubectl apply -f ./k8s/ns.yaml
	kubectl apply -f ./k8s/

undeploy:
	kubectl delete -f ./k8s/
	# kubectl delete namespace kubeapi-apiserver &
	# sleep 5
	# kubectl get namespace kubeapi-apiserver -o json > ./logging.json
	# grep -F -v \"kubernetes\" logging.json > logging.json.tmp && mv logging.json.tmp logging.json
	# kubectl replace --raw "/api/v1/namespaces/kubeapi-apiserver/finalize" -f ./logging.json
	# rm -rf ./logging.json

run:
	sudo env "PATH=${PATH}" go run . --etcd-servers localhost:2379 \
    --authentication-kubeconfig ${HOME}/.kube/config \
    --authorization-kubeconfig ${HOME}/.kube/config \
    --kubeconfig ${HOME}/.kube/config