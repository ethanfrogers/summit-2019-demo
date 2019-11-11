FROM ubuntu:18.04

RUN apt-get update && apt-get install -y curl

ADD fetch /usr/local/bin

CMD ["/usr/local/bin/fetch"]BB