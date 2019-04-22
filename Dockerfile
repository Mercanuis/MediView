FROM golang:1.12

LABEL author="Joseph Orme <brogramn@gmail.com>"

WORKDIR $GOPATH/src/github.com/MediView

COPY . ./

RUN go install -v ./...

RUN go build ./...

RUN \
	make build && \
	cp bin/mediview /go/bin/mediview

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/mediview /bin/mediview

CMD ["mediview"]