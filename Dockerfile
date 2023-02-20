FROM alpine:3.17

WORKDIR /app

HEALTHCHECK --interval=30s --timeout=5s \
    CMD wget -q -O - http://$HOSTNAME:17000/health-check || exit 1

COPY . .

ENTRYPOINT [ "bin/mono" ]

CMD [ "serve" ]
