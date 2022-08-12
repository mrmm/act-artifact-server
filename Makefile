docker-build:
	docker build -t elron/act-artifact-server:latest .

docker-push:
	docker push elron/act-artifact-server:latest

build-push:
	docker-build
	docker-push