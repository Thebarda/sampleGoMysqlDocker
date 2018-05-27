FROM golang:1.10.2

ARG app_env
ENV APP_ENV $app_env

COPY ./ /backend/app
WORKDIR /backend/app

RUN go get -d ./...
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	./app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
  fi
EXPOSE 8081