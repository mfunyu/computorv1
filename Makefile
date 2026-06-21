
all		:
	go build

format :
	gofmt -s -w .
	goimports