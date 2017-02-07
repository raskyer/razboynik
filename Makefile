build:
	go get
	go build

install:
	go install

clean: 
	rm pkg/pkg
	rm plugin/bin/*
