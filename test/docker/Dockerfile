FROM ubuntu:focal

RUN apt-get update --fix-missing
RUN apt-get -y install git curl bash

RUN mkdir -p /app
WORKDIR /app

RUN curl -L https://go.dev/dl/go1.19.1.linux-amd64.tar.gz -o go.tar.gz
RUN tar -xf go.tar.gz

ENV PATH="/app/go/bin:$PATH"
ENV GOPATH="/app/godata"
ENV GOROOT="/app/go"

ARG CACHEBUST=1
RUN go install github.com/longht021189/ops-cl@latest

CMD ["bash"]