FROM golang:latest
ENV REPO_URL github.com/humbertodias/go-nie-api
RUN go get ${REPO_URL} && go install ${REPO_URL}
CMD go-nie-api
EXPOSE 8080