GO := $(shell which go)
GOINSTALL=$(GO) install

EXENAME=client1/main

export GOPATH := $(shell pwd)
export GOBIN := $(shell pwd)/bin

all: dep build docker-build

dep:
	@echo Resolving dependencies...	
	@cd src/client1 && dep ensure

build:
	@echo Build Started...	
	@$(GOINSTALL) $(EXENAME)
	@if [ $$? != 0 ] ; then echo "Build failed" exit 1 ; fi
	@echo Build Success..

docker-build:
	@echo Building docker image...
	@docker build -t grpc_c1ient1 .
	@echo Image built successfully
