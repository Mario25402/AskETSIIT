FROM alpine:latest 

RUN apk update \
    && apk add make \
    && wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz \
    && rm go1.23.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir -p /app/test

WORKDIR /app/test

COPY . /app/test/

RUN make install

RUN adduser -D demo

USER demo

ENTRYPOINT [ "make", "test" ]