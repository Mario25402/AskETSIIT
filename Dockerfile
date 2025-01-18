FROM alpine:latest 

WORKDIR /app/test

RUN apk update \
    && apk add make \
    && wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz \
    && rm go1.23.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOCACHE=/app/test/.cache

RUN mkdir -p /app/test/.cache

ARG USER_ID
RUN if [ -z "$USER_ID" ]; then USER_ID=1001; fi \
    && adduser -D -u $USER_ID dynamicuser \
    && chown -R $USER_ID:$USER_ID /app/test

USER dynamicuser

ENTRYPOINT [ "make", "test" ]