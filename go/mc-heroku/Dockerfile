FROM openjdk:8u212-jre-alpine

RUN apk -U --no-cache upgrade

RUN apk add --no-cache -U \
    vim \
    bash

RUN mkdir /minecraft
COPY mc-heroku /minecraft/mc-heroku
COPY token.json /minecraft/token.json
COPY entrypoint.sh /minecraft/entrypoint.sh

RUN chmod 777 mc-heroku
RUN chmod 444 token.json
RUN chmod +x /entrypoint.sh

# ENTRYPOINT ["/entrypoint.sh"]

