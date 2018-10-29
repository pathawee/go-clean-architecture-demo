FROM golang:1.9.4-alpine3.7
RUN apk update && apk add git
ARG app_env
ENV APP_ENV $app_env

RUN go get github.com/golang/dep/cmd/dep
WORKDIR /go/src/go-clean-architecture-demo
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock
RUN dep ensure --vendor-only
ADD app app

CMD if [ ${APP_ENV} != development ]; \
	then \
	cd app/ && go build && ./app; \
	else \
	go get github.com/pilu/fresh && cd app/ && fresh; \
	fi
