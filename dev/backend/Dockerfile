FROM golang:1.9-alpine

ARG packagename
ENV MOUNT_DIR /go/src/$packagename/

RUN mkdir -p $MOUNT_DIR /build/watchman

# install watchman
ADD https://github.com/facebook/watchman/archive/v4.7.0.zip     /build/watchman

RUN apk add --update --no-cache python python-dev py-pip make gcc g++ automake autoconf linux-headers
RUN cd /build/watchman && \
    unzip v4.7.0.zip && \
    cd watchman-4.7.0 && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install

RUN pip install pywatchman

RUN apk del python-dev py-pip automake autoconf linux-headers
# per watchman adjust /proc/sys/fs/inotify/max* as needed - trying defaults first

RUN ln -s $MOUNT_DIR /backend
VOLUME /backend

COPY watch_and_run.sh /entrypoint.sh
COPY run.sh /run.sh

RUN chmod +x /run.sh && chmod +x /entrypoint.sh

WORKDIR $MOUNT_DIR
ENTRYPOINT ["/entrypoint.sh"]

