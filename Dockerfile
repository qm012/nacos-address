FROM alpine:latest

MAINTAINER urobot@qq.com

ENV VERSION="1.1.0"

WORKDIR /home/nacos-address/

RUN wget https://github.com/qm012/nacos-address/releases/download/v${VERSION}/nacos-address-${VERSION}.tar.gz -P /home

RUN tar -zxvf /home/nacos-address-${VERSION}.tar.gz -C /home \
    && rm -rf /home/nacos-address-${VERSION}.tar.gz

VOLUME /home/nacos-address/config
VOLUME /home/nacos-address/logs

EXPOSE 8849

CMD ["./nacos-address"]