FROM alpine:latest 

WORKDIR /app/test

RUN apk update \
    && apk add --no-cache make curl \
    && GO_VERSION=$(curl -Ls https://go.dev/VERSION?m=text | head -n 1) \
    && wget https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf ${GO_VERSION}.linux-amd64.tar.gz \
    && rm ${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOCACHE=/home/tests/.cache

RUN adduser -D -h /home/tests tests \
    && mkdir -p /home/tests/.cache \
    && chmod -R a+w /home/tests/.cache

USER tests

ENTRYPOINT [ "make", "test" ]