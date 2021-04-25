FROM golang AS builder
ENV GOOS=linux
ENV GOARCH=386
WORKDIR /src
COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY go.mod .
COPY go.sum .
## Compiling *.go file
RUN ls -la
RUN go build -a cmd/main.go

FROM scratch
LABEL AUTHOR=kslisenka
EXPOSE 9090
WORKDIR /app
COPY --from=builder /src/main .
## Define container process
CMD ["./main"]