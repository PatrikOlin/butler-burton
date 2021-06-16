FROM golang:1.16-alpine3.12 as build

WORKDIR /app
COPY . /app

RUN go build -o bb *.go

FROM alpine:3.13.5
ENV USER="Butler Burton"

WORKDIR /app
COPY --from=build /app/bb /app 

COPY certs/private.json /root/.butlerburton/private.json
COPY certs/butlerBurtonCert.pfx /root/.butlerburton/butlerBurtonCert.pfx

RUN cp bb /usr/local/bin/bb

CMD [ "/bin/sh" ]
