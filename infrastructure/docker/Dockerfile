FROM golang:1.21

WORKDIR /fase-4-hf-voucher/src/app

COPY . .

ENV GOPATH=/fase-4-hf-voucher

#RUN go build -ldflags "-w -s" -o bin/hermesfoods cmd/app/*.go
RUN go build -ldflags "-w -s" -o bin/hermesfoods cmd/app/main.go


ENTRYPOINT ["./bin/hermesfoods"]

EXPOSE 8082
