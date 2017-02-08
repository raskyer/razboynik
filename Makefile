build:
	cd pkg && go build -o ../razboynik
	make plugin

install:
	cd pkg && go get
	cd pkg && go install

plugin:
	cd plugin && ./build_all.sh shellmodule
	cd plugin && ./build_all.sh phpmodule
	cd plugin && ./build_all.sh kernelmodule

clean: 
	rm razboynik
	rm plugin/bin/*

clean_pkg:
	rm pkg/pkg

clean_plugin:
	rm plugin/bin/*
