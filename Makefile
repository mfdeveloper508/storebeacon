ROOT_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	docker build -t fupp-api ./

build-production:
	docker build -t fupp-api ./ --build-arg app_env=production

serve:
	docker run -it -p 8080:8080 --net="host" -v ${ROOT_DIR}:/go/src/github.com/store fupp-api

test:
	docker run --entrypoint "run-tests.sh" fupp-api
