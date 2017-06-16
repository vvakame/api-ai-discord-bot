# use Docker multi stage build

FROM golang:1.8.3 AS build-env

ADD . /go/src/github.com/vvakame/api-ai-discord-bot
WORKDIR /go/src/github.com/vvakame/api-ai-discord-bot
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o bridge-bot -a -tags netgo -installsuffix netgo bot/bot.go


FROM alpine

COPY --from=build-env /go/src/github.com/vvakame/api-ai-discord-bot/bridge-bot /usr/local/bin/bridge-bot
RUN apk add --no-cache --update ca-certificates

ENTRYPOINT ["/usr/local/bin/bridge-bot"]

EXPOSE 8080
