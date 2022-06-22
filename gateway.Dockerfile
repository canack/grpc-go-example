FROM golang AS go-build-env

COPY services /services
WORKDIR /services/gateway

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

FROM gcr.io/distroless/static
COPY --from=go-build-env /services/gateway/gateway-service /

CMD ["/gateway-service"]