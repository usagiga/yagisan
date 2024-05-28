FROM golang:1.22.3 as builder
WORKDIR /src
COPY . .
RUN make build

FROM gcr.io/distroless/static-debian12
COPY --from=builder /src/yagisan .
ENTRYPOINT [ "./yagisan" ]
