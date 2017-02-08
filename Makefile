build:
	cd pkg && go get
	cd pkg && go build

install:
	cd pkg && go install

plugin:
	cd plugin && ./build_all.sh shellmodule
	cd plugin && ./build_all.sh phpmodule
	cd plugin && ./build_all.sh kernelmodule

clean: 
	rm pkg/pkg
	rm plugin/bin/*

clean_pkg:
	rm pkg/pkg

clean_plugin:
	rm plugin/bin/*
