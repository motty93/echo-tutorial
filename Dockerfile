ARG GO_VERSION=1.14.6
ARG ALPINE_VERSION=3.12

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ENV APP_ROOT /go/src/app
ENV GO111MODULE=on

WORKDIR ${APP_ROOT}
COPY ./src .

RUN apk update \
    && apk add --no-cache git \
    && git config --global http.postBuffer 524288000 \
    && rm -rf /var/cache/apk/*

# air install
RUN go build -o app \
    && go get -u github.com/cosmtrek/air

EXPOSE 8020
CMD ["air"]
