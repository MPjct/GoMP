

test:
	go test ./... -cover
        
bench:
	go test ./... -cover -bench .

fmt:
	find . -name \*.go -type f -exec go fmt {} \;

format:
	make fmt

