FROM alpine:edge

ADD . /srv/src/github.com/bodokaiser/retina

# gcc fails at opencv with an internal compiler error
ENV CC /usr/bin/clang
ENV CXX /usr/bin/clang++

ENV GOPATH /srv
ENV CVPATH /srv/opencv

ENV PKG_CONFIG_PATH /usr/local/lib/pkgconfig

RUN apk add -U go git pkgconf build-base cmake clang linux-headers
RUN git clone --depth 1 http://github.com/itseez/opencv $CVPATH

WORKDIR $CVPATH

RUN cmake \
    -D CMAKE_BUILD_TYPE=RELEASE \
    -D CMAKE_INSTALL_PREFIX=/usr/local \
    -D BUILD_TESTS=OFF \
    -D BUILD_EXAMPLES=OFF \
    -D BUILD_PERF_TESTS=OFF \
    -D WITH_OPENGL=OFF \
    .

RUN make && make install && rm -rf $CVPATH

# workaround to linking problem by using gcc over clang
RUN CC="" CXX="" go get github.com/bodokaiser/retina/...
RUN go build github.com/bodokaiser/retina/cmd/http

CMD ["/srv/bin/http"]

EXPOSE 3000
