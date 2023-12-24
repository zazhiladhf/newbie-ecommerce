FROM golang as builder

WORKDIR /apps

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o main main.go

FROM alpine

WORKDIR /apps

COPY --from=builder /apps/main . 

CMD [ "./main" ]