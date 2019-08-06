FROM golang:1.9 AS builder

RUN go version

COPY . "/go/src"
WORKDIR "/go/src"

#RUN go get -v -t  .
RUN set -x && \
    go get github.com/gorilla/mux && \
    go get github.com/suyashkumar/dicom

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o huffmanEncoder

CMD ["./huffmanEncoder"]

EXPOSE 8080


FROM scratch

COPY --from=builder /go/src/huffmanEncoder .

EXPOSE 8080

CMD ["./huffmanEncoder"]