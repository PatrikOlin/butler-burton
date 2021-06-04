FROM golang:1.16-alpine3.12 as build

WORKDIR /app
COPY . /app

RUN go build -o bb *.go

FROM alpine:3.13.5

WORKDIR /app
COPY --from=build /app/bb /app 
COPY --from=build /app/assets/config.yml /root/.config/butlerburton/config.yml
COPY --from=build /app/assets/report.xlsx /root/.butlerburton/report.xlsx

RUN cp bb /usr/local/bin/bb

RUN [ "./bb", "report", "set", "report.xlsx" ]
CMD [ "/bin/sh" ]
