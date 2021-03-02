FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /Users/shengchao/Program/golang/gin-blog
COPY . /Users/shengchao/Program/golang/gin-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]