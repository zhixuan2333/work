FROM openjdk:8u212

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

RUN curl -Lo /usr/local/bin/frpc https://getfrp.sh/d/frpc_linux_amd64 \
 && chmod 755 /usr/local/bin/frpc

ENTRYPOINT ["/entrypoint.sh"]
