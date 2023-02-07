FROM golang
ADD . /go/src/myapp
RUN go get gopkg.in/mgo.v2  
RUN go get gopkg.in/mgo.v2/bson
RUN go get -u github.com/gorilla/mux
RUN go install myapp
ENTRYPOINT /go/bin/myapp
EXPOSE 8088