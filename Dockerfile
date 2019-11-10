FROM golang:1.13.4-alpine3.10 as builder

ENV GO111MODULE=onB GOOS=linux GOARCH=amd64

WORKDIR /opt/armory/build/
ADD ./ /opt/armory/build/

RUN go mod download && go build

FROM alpine:3.10

COPY --from=builder /opt/armory/build/summit-2019-demo /opt/armory/

CMD ["/opt/armory/summit-2019-demo"]