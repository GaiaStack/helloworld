FROM golang:1.4.2

ADD . $GOPATH/src

ADD run.sh /run.sh

RUN chmod +x /run.sh

EXPOSE 80

WORKDIR $GOPATH/src

CMD ["/run.sh"]
