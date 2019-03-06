FROM golang:alpine as builder

RUN mkdir /user && \
    echo 'user:x:504:504:user:/home/user:' > /user/passwd && \
    echo 'user:x:504:user' > /user/group

WORKDIR /go/src
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix 'static' -o /fwew .

FROM scratch

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /fwew /fwew
COPY --from=builder /go/src/.fwew/config.json /go/src/.fwew/dictionary.txt /home/user/.fwew/

USER user:user
ENTRYPOINT ["/fwew"]
