FROM golang:1.19-bullseye as builder

ADD . /go/crssy
WORKDIR /go/crssy
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/crssy/crssy /usr/bin/crssy
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/crssy" ]