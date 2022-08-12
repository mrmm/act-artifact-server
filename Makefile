docker-build:
	docker build -t elron/act-artifact-server:latest .

docker-push:
	docker push elron/act-artifact-server:latest

docker-run-image:
	docker run --rm -p 1234:1234 elron/act-artifact-server:latest -p 1234

build-push:
	docker-build
	docker-push