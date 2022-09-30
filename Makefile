CURRENT_DIR = $(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

gen-proto:
	./scripts/gen-proto.sh $(CURRENT_DIR)