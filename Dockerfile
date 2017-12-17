FROM alpine
LABEL maintainer="drmfly.liw@gmail.com"
WORKDIR /srv/gopkg
ADD qtraderpkg ./
ENTRYPOINT ["./qtraderpkg"]