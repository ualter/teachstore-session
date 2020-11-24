FROM golang

RUN mkdir -p /go/src/teachstore-session
ADD . /go/src/teachstore-session
WORKDIR /go

RUN go get ./...
RUN go install -v ./...

CMD ["/go/bin/teachstore"]