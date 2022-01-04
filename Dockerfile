FROM registry-in.dustess.com:9000/base/centos:7
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone
WORKDIR /app
ADD ./work-order-console /app
ENV RUN_MODE pro
ENV GIN_MODE release
EXPOSE 8080 8080
CMD ["/app/work-order-console"]
