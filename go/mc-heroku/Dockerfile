FROM openjdk:8u212

RUN apt-get update \
 && DEBIAN_FRONTEND="noninteractive" apt-get install -y --no-install-recommends \
    vim \
    bash \
    git \
    openssh-client \
 && apt-get -y clean \
 && rm -rf /var/lib/apt/lists/*

RUN curl -Lo /usr/local/bin/frpc https://getfrp.sh/d/frpc_linux_amd64 \
 && curl -Lo /minecraft/app/server.jar https://launcher.mojang.com/v1/objects/f02f4473dbf152c23d7d484952121db0b36698cb/server.jar \
 && chmod 555 /usr/local/bin/frpc \
 && chmod 555 /minecraft/app/server.jar

# Add the keys and set permissions
COPY id_rsa /root/.ssh/id_rsa

RUN mkdir /minecraft
RUN mkdir /minecraft/app
COPY mc-heroku /minecraft/mc-heroku
COPY entrypoint.sh /minecraft/entrypoint.sh
COPY sync.sh /minecraft/sync.sh

RUN chmod 555 /minecraft/mc-heroku \
 && chmod 555 /root/.ssh/id_rsa \
 && chmod +x /minecraft/entrypoint.sh \
 && chmod +x /minecraft/sync.sh

WORKDIR /minecraft/app
CMD /minecraft/entrypoint.sh
