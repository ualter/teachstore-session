FROM golang

RUN mkdir -p /go/src/teachstore-session
ADD . /go/src/teachstore-session
WORKDIR /go/src/teachstore-session

RUN go get -d ./...
RUN go install -v ./...

RUN rm -rf /go/src/teachstore-session
RUN mkdir -p /go/src/teachstore-session/config
ADD ./config /go/src/teachstore-session/config

CMD ["/go/bin/teachstore-session"]