FROM openjdk:8u212

RUN apk -U --no-cache upgrade

RUN apk add --no-cache -U \
    vim \
    bash

RUN mkdir /minecraft
COPY mc-heroku /minecraft/mc-heroku
COPY token.json /minecraft/token.json
COPY entrypoint.sh /minecraft/entrypoint.sh

RUN chmod 777 /minecraft/mc-heroku
RUN chmod 444 /minecraft/token.json
RUN chmod +x /minecraft/entrypoint.sh

# ENTRYPOINT ["/entrypoint.sh"]
