# docker user name
NAMESPACE := sarodesparsh

all: notification-server-docker notification-worker-docker

golang-base-docker:
	docker build -t $(NAMESPACE)/golang-base -f ./golang-base/Dockerfile ../

notification-server-docker: golang-base-docker
	docker build -t $(NAMESPACE)/notification-server -f ./notification-server/Dockerfile ../ --build-arg="NAMESPACE=$(NAMESPACE)"
	docker push $(NAMESPACE)/notification-server

notification-worker-docker:
	docker build -t $(NAMESPACE)/notification-worker -f ./notification-worker/Dockerfile ../
	docker push $(NAMESPACE)/notification-worker

.PHONY: all golang-base-docker notification-server-docker notification-worker-docker
