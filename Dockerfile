FROM golang:1.16-alpine3.12 as build

WORKDIR /app
COPY . /app

RUN go build -o bb *.go

FROM alpine:3.13.5

WORKDIR /app
COPY --from=build /app/bb /app 
COPY --from=build /app/assets/config.yml /root/.config/butlerburton/config.yml
COPY --from=build /app/assets/report.xlsx /root/.butlerburton/report.xlsx

COPY certs/private.json /root/.config/butlerburton/private.json
COPY certs/butlerBurtonCert.pfx /root/.config/butlerburton/butlerBurtonCert.pfx

RUN cp bb /usr/local/bin/bb

CMD [ "/bin/sh" ]
