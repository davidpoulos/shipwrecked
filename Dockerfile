FROM golang:1.14.3-stretch

COPY . /go
RUN unset GOPATH

CMD ["sleep", "infinity"]