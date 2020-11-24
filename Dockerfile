FROM golang

RUN mkdir -p /go/src/teachstore-session
ADD . /go/src/teachstore-session
WORKDIR /go/src/teachstore-session

RUN go get -d ./...
RUN go install -v ./...

CMD ["/go/bin/teachstore-session"]