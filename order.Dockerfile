FROM golang AS go-build-env

COPY services /services
WORKDIR /services/order

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

FROM gcr.io/distroless/static
COPY --from=go-build-env /services/order/order-service /
COPY --from=go-build-env /services/.aws/config /root/.aws/
COPY --from=go-build-env /services/.aws/credentials /root/.aws/

CMD ["/order-service"]