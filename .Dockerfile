FROM golang:alpine AS builder

WORKDIR /app
COPY . /app/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/url-shortner cmd/url-shortner/main.go


FROM scratch

WORKDIR /app
COPY --from=builder /app/url-shortner /app/

EXPOSE 8080

CMD [ "/app/url-shortner" ]