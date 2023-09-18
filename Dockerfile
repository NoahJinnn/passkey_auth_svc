FROM alpine:3.17

WORKDIR /app

COPY /bin .

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
    echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
    apk add doppler

CMD [ "doppler", "run", "--", "./hq", "serve", "--wa.id", "$AUTH_DOMAIN", "--wa.origins", "https://${AUTH_DOMAIN}" ]

