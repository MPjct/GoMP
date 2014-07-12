

test:
	go test ./... -cover
        
bench:
	go test ./... -cover -bench .

fmt:
	find . -name \*.go -type f -exec go fmt {} \;

format:
	make fmt

cover:
	go test github.com/MPjct/GoMP/MySQLProtocol -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
