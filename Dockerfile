FROM golang:alpine

WORKDIR /go/src
COPY . .

RUN adduser -D user
RUN apk add --no-cache git make
RUN go get -d -v ./...
RUN make
RUN cp bin/fwew /go/bin/
RUN cp -r .fwew /home/user/

USER user
ENTRYPOINT ["fwew"]
