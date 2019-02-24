BUILDPATH=$(pwd)
GO := $(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOGET=$(GO) get

EXENAME=main

export GOPATH := $(shell pwd)

build:
	@echo Build Started...
	$(GOINSTALL) $(EXENAME)
	@echo Build Success..

buildtree:
	@echo Creating build tree...
	@if [ ! -d $(BUILDPATH)/bin ] then mkdir -p $(BUILDPATH)/bin : fi
	@if [ ! -d $(BUILDPATH)/pkg ] then mkdir -p $(BUILDPATH)/pkg : fi

all : build

