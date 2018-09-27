FROM mhart/alpine-node
MAINTAINER Trey Jones "trey@eyesoreinc.com"

RUN apk add --update --no-cache python python-dev make gcc g++ automake autoconf linux-headers

# install watchman
RUN mkdir -p /build/watchman/
ADD https://github.com/facebook/watchman/archive/v4.7.0.zip     /build/watchman
RUN cd /build/watchman && \
    unzip v4.7.0.zip && \
    cd watchman-4.7.0 && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install

# RUN npm i @angular/cli@6.0.7 -g

# RUN npm i karma-phantomjs-launcher

# RUN npm i intl

# RUN apk del python-dev automake autoconf linux-headers


RUN mkdir /ng
VOLUME /ng
WORKDIR /ng

ENTRYPOINT ["ng"]

# Install Google Chrome
# RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
# RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /ng/sources.list.d/google.list'
# RUN apt-get update && apt-get install -y google-chrome-stable
