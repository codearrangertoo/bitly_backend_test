FROM golang:1.15 AS build
WORKDIR /usr/local/src/bitly_backend_test/
COPY ./ .
RUN CGO_ENABLED=0 go build .

FROM scratch AS final
WORKDIR /
COPY --from=build /usr/local/src/bitly_backend_test/bitly_backend_test /bin/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080/tcp
ENTRYPOINT ["/bin/bitly_backend_test"]
