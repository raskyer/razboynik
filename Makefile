build:
	cd pkg && go get
	cd pkg && go build

install:
	cd pkg && go install

plugin:
	cd plugin && ./build_all.sh

clean: 
	rm pkg/pkg
	rm plugin/bin/*

clean_plugin:
	rm plugin/bin/*
