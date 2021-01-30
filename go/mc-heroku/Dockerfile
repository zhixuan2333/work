FROM openjdk:8u212-jre-alpine

RUN apk -U --no-cache upgrade

RUN apk add --no-cache -U \
    vim \
    bash

COPY mc-heroku mc-heroku
COPY token.json token.json
COPY entrypoint.sh entrypoint.sh

RUN chmod 777 mc-heroku
RUN chmod 444 token.json
RUN chmod +x /entrypoint.sh

# ENTRYPOINT ["/entrypoint.sh"]

