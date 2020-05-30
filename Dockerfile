FROM golang:1.14.3-stretch

COPY . /go


CMD ["sleep", "infiniti"]