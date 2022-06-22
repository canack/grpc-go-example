FROM golang AS go-build-env

COPY services /services
WORKDIR /services/customer

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

FROM gcr.io/distroless/static
COPY --from=go-build-env /services/customer/customer-service /
COPY --from=go-build-env /services/.aws/config /root/.aws/
COPY --from=go-build-env /services/.aws/credentials /root/.aws/

CMD ["/customer-service"]

