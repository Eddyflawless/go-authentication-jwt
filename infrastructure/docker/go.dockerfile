ARG ALPINE_VERION=1.13-alpine2.11
FROM  golang:${ALPINE_VERION} as builder

RUN apk --no-cache --update add \
    autoconf \
    automake \
    bash \
    build-base \
    font-noto-emoji \
    g++ \
    gcc \
    gcompat \
    libstdc++ \
    libtool \
    make \
    nasm \
    py-pip \
    python3 \
    rsync \
    curl \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/lib/apt/lists/*


RUN mkdir /build
RUN  cp -r ./api/* go.mod go.sum /build

WORKDIR /build
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main .

ENTRYPOINT [ "./main" ]