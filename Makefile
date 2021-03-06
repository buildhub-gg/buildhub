MAIN_SOURCE_SETS=./server.go ./wire_gen.go

run:
	go run $(MAIN_SOURCE_SETS)
build:
	go run $(MAIN_SOURCE_SETS)
generate:
	wire
	go generate ./...
