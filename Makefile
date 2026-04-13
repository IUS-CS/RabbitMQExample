all: send/send recv/recv

serve-rabbitmq:
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management

serve-send: send/send
	./send/send

serve-recv: recv/recv
	./recv/recv

testmessage:
	curl -X POST -d '{"name": "John", "email": "foo@bar.com"}' 127.0.0.1:8080

person/person.pb.go: person.proto
	protoc --go_out=./ person.proto

send/send: send/send.go person/person.pb.go common/common.go
	go build -o send send/send.go

recv/recv: recv/recv.go person/person.pb.go common/common.go
	go build -o recv recv/recv.go

rmq.pdf: rmq.typ
	typst c rmq.typ

clean:
	@rm -rf person recv/recv send/send rmq.pdf

.PHONY: all serve-rabbitmq serve-recv serve-send testmessage clean
