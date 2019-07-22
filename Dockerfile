FROM golang:1.12.6-stretch

RUN export PATH=$PATH:$GOPATH/bin

RUN go get github.com/99designs/gqlgen
RUN go get github.com/cespare/reflex
RUN go get github.com/mattdamon108/gqlmerge

COPY docker/go.mod /app/
COPY docker/go.sum /app/

COPY reflex.conf /reflex/

COPY docker/docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]