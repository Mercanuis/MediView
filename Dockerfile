FROM golang:1.12

LABEL author="Joseph Orme <brogramn@gmail.com>"

WORKDIR $GOPATH/src/github.com/MediView

COPY . .

RUN echo "starting installation"
RUN go get -d -v ./...
RUN go install -v ./...
RUN echo "installtion complete"

EXPOSE 20001

CMD ["bin/mediview"]