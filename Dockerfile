FROM golang:1.12.6-stretch

RUN export PATH=$PATH:$GOPATH/bin

RUN go get github.com/99designs/gqlgen
RUN go get github.com/cespare/reflex

COPY docker/go.mod /app/
COPY docker/go.sum /app/
COPY reflex.conf /reflex/

COPY docker/docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]