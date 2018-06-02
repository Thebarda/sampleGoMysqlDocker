FROM golang:1.10.2

ARG app_env
ENV APP_ENV $app_env

COPY ./ /backend/app
WORKDIR /backend/app

RUN go get -d ./...
RUN go get github.com/stretchr/testify/assert
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	./app; \
	else  \
	if [ ${APP_ENV} = test ]; \
	then \
	go test ./tests/ -v; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi \
  fi
EXPOSE 8081