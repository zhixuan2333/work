FROM openjdk:8u212

RUN apt -U --no-cache upgrade

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    vim \
    bash \
 && apt-get -y clean \
 && rm -rf /var/lib/apt/lists/*

RUN mkdir /minecraft
COPY mc-heroku /minecraft/mc-heroku
COPY token.json /minecraft/token.json
COPY entrypoint.sh /minecraft/entrypoint.sh

RUN chmod 777 /minecraft/mc-heroku
RUN chmod 444 /minecraft/token.json
RUN chmod +x /minecraft/entrypoint.sh

# ENTRYPOINT ["/entrypoint.sh"]
