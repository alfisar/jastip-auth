FROM golang:1.20.0-buster

RUN mkdir -p /go/src/jastip

WORKDIR /go/src/jastip
COPY . .

RUN apt update
RUN apt install -y tzdata
ENV GOPRIVATE="github.com/alfisar"
ENV TZ Asia/Jakarta

RUN git config --global \
    url."https://${git_username}:${git_token}@github.com/".insteadOf \
    "https://github.com"

RUN go get -d -v ./...

RUN go build -o /go/bin/jastip

RUN rm -rf /go/src/jastip/.git
RUN rm -rf $HOME/.gitconfig

EXPOSE 8801

CMD ["jastip"]