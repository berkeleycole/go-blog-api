FROM golang:1.20.3 AS builder

WORKDIR /app
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /bin/blog-api

EXPOSE 80

CMD ["/bin/blog-api"]


# PRODUCTION

FROM debian AS production

COPY --from=builder /bin/blog.api /bin

CMD ["/bin/blog-api"]