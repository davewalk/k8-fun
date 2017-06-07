version:=$(shell cat main.go | sed -rn 's/version = "(.+)"/\1 /p' | xargs)
project_id:=propane-folio-103015
container_name:=sample-server-app

build:
		CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s"
		docker build -t gcr.io/$(project_id)/sample-server:$(version) .		
run:
		docker run -it --rm -p 8080:8080 -e "PORT=8080" --name $(container_name) gcr.io/$(project_id)/sample-server:$(version) 
push:
		gcloud docker -- push gcr.io/$(project_id)/sample-server:$(version)
apply:
		kubectl apply -f k8-deployment.yml
create:
		kubectl create -f k8-deployment.yml
bash:
		container_id:=$(docker ps -aqf "name=$(container_name)"
		docker exec it $(container_id) /bin/bash
expose:
		rs:=$(kubectl get rs --selector=app=sample-server --output=json | jq '.items[] | select(.status.availableReplicas > 0) | .metadata.name')
		kubectl expose rs $(rs) --type="LoadBalancer" --name="sample-service"
