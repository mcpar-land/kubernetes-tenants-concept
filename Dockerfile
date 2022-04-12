FROM golang:alpine
COPY ./helloenv /app
WORKDIR /app
RUN go build -o /app/helloenv .

FROM alpine
COPY --from=0 /app/helloenv /bin/helloenv
CMD ["helloenv"]