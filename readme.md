# RabbitMQ Example

This set of example programs runs a RabbitMQ server, a receiving server that consumes a queue, and a sending server that sends to the queue from an HTTP endpoint.

## Requirements

* Docker
* Go
* Protobuf (with the Go protobuf tools)
* Make/bash for ease of running
* curl for the testmessage target

## Building

All steps are in the Makefile.

* `make serve-rabbitmq` starts the RabbitMQ server
* `make serve-recv` starts the receiving server
* `make serve-send` starts the sending server
* `make testmessage` sends a message through the system

You can, of course, run all of these make steps yourself or modify them. The makefile exists as a reference for the commands used in class.
