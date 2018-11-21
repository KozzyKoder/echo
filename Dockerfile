FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
RUN adduser -S -D -H -h /app appuser
WORKDIR /app
RUN ["chmod", "+x", "./main"]
USER appuser
ENTRYPOINT ["./main"]