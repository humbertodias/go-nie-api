FROM golang:latest
ENV WKDIR /go/src/github.com/humbertodias/go-nie-api
RUN mkdir -p ${WKDIR}
ADD . ${WKDIR}
EXPOSE 8080
WORKDIR ${WKDIR}
RUN make get
CMD make run