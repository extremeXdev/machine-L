FROM arm64v8/golang:1.20.6-alpine3.17

LABEL maintainer="EYNDO <eyndo.net>"

RUN apk add --no-cache ca-certificates

USER root

COPY . /go/src/github.com/extremeXdev/machine-L/Projet/api

EXPOSE 2501
WORKDIR /go/src/github.com/extremeXdev/machine-L/Projet/api
RUN chmod +x run.sh
CMD ["./run.sh"]

