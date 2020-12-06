FROM golang

RUN mkdir -p /go/src/teachstore-session
ADD . /go/src/teachstore-session
WORKDIR /go/src/teachstore-session

RUN go get -d ./...
RUN go install -v ./...

RUN rm -rf /go/src/teachstore-session
RUN mkdir -p /go/src/teachstore-session/config

# Add this line to run without K8s Config (Local docker only) - Running with K8s ConfigMap this line is not necessary
# ADD ./config /go/src/teachstore-session/config

ENV GOTRACEBACK=single
CMD ["/go/bin/teachstore-session"]