FROM golang:1.19.5

EXPOSE "8888"

WORKDIR ./

RUN ["go","build"]

CMD ["./dousheng"]

