FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get github.com/joseabad123/SimpleGoApp/src
RUN cd /build && git clone https://github.com/joseabad123/SimpleGoApp.git

RUN cd /build/SimpleGoApp/src && go build

EXPOSE 5656

ENTRYPOINT [ "/build/SimpleGoApp/src/main" ]