FROM alpine:latest 

WORKDIR /app/test

RUN apk update \
    && apk add make \
    && wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz \
    && rm go1.23.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOCACHE=/home/tests/.cache

RUN adduser -D -h /home/tests tests \
    && mkdir -p /home/tests/.cache \
    && chmod -R a+w /home/tests/.cache

USER tests

ENTRYPOINT [ "make", "test" ]