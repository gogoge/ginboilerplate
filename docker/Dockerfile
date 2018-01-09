FROM golang:latest 
RUN go get -u -v github.com/kardianos/govendor && \ 
    mkdir -p /go/src/devpack.cc/webb_lu/ginboilerplate 
COPY . /go/src/devpack.cc/webb_lu/ginboilerplate/ 
WORKDIR /go/src/devpack.cc/webb_lu/ginboilerplate 
RUN govendor sync && \
    go build -o main . 
CMD ["/go/src/devpack.cc/webb_lu/ginboilerplate/main"]