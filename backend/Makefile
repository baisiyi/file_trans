.PHONY: proto build
proto:
	./scripts/gen-proto.sh

build:
	go build -o file_server

run:
	go build -o file_server && chmod +x file_server && ./file_server -config=config/config.yaml

clean:
	rm -f file_server