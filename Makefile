build:
	go build .

image:
	docker build -t hello-world:1.0 -f Dockerfile

clean:
	go clean
