FROM golang:1.7.4
LABEL MAINTAINER="dad264@psu.edu"

ENV SOURCES /go/src/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "microservice" ]

