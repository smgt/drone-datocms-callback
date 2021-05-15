FROM golang:1.16 as builder

WORKDIR /release

COPY /go.mod /go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -v -o datocms

FROM scratch
#ENV GODEBUG netdns=go

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /release/datocms /

CMD ["/datocms"]
