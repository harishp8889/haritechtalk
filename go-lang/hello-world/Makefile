BUILDPATH=$(pwd)
GO := $(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH := $(shell pwd)

all: buildtree build docker-build

buildtree:
	@echo Creating build tree...
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

build:
	@echo Build Started...
	@$(GOINSTALL) $(EXENAME)
	@if [ $$? != 0 ] ; then echo "Build failed" exit 1 ; fi
	@echo Build Success..

docker-build:
	docker build -t go-hello .
	@echo Image built successfully
	
