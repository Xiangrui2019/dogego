default: deps
	go install

deps:
	go mod tidy
	go mod vendor

run: docker
	sudo docker-compose up

docker: deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	chmod +x ./main
	sudo docker build -t go-shop-pay .

clean:
	rm -rf vendor/
	rm -rf go.sum

.PHONY: go test