graph:
	go run core/flow/graph/main.go | dot -T png -o core/flow/graph/graph.png -Ksfdp

enums:
	go generate core/flow/*.go

build-server:
	go build -o bin/geralt-server main.go

build-client:
	go build -o bin/geralt client/main.go

build: build-client build-server
	echo "Built client and server"