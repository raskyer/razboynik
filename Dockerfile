FROM golang:1.6-wheezy

RUN apt-get update && apt-get install git

ADD . /go/src/github.com/eatbytes/razboynik/

RUN cd /go/src/github.com/eatbytes/ && git clone https://github.com/EatBytes/razboy.git
RUN cd /go/src/github.com/eatbytes/razboy && git checkout release/2.0.0

RUN cd /go/src/github.com/eatbytes/razboynik && go get
RUN go install github.com/eatbytes/razboynik
