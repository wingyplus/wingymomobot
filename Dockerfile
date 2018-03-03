FROM golang:alpine AS builder

COPY . /go/src/github.com/wingyplus/wingymomobot/
RUN go install github.com/wingyplus/wingymomobot

FROM alpine

RUN apk update && apk add ca-certificates
RUN adduser -D wingymomo
USER wingymomo

COPY --from=builder /go/bin/wingymomobot /usr/local/bin/
CMD ["/usr/local/bin/wingymomobot"]
