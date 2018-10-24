FROM golang:1.9.4-alpine3.7
RUN apk update && apk add git
# install dep
RUN go get github.com/golang/dep/cmd/dep
# install gin
RUN go get github.com/codegangsta/gin
# create a working directory
WORKDIR /go/src/go-clean-architecture-demo
# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock
# install packages
RUN dep ensure --vendor-only
# add source code app to src/go-clean-architecture-demo folder
ADD app app

CMD ["go", "run", "app/main.go"]

#CMD ["go", "build", "app/main.go"]
#CMD ["./main"]
