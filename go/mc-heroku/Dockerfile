FROM ubuntu:20.04

RUN apt-get update \ 
    && apt-get install -y locales \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y \
    vim \
    curl \
    git \
    ssh \
    sudo \
    vim 

RUN apt-get install -y wget 

RUN cd /tmp \
    && wget https://github.com/cdr/code-server/releases/download/v3.7.2/code-server-3.7.2-linux-amd64.tar.gz 

RUN cd /tmp \
    && tar xzf code-server-3.7.2-linux-amd64.tar.gz \
    && mkdir /usr/local/lib/code-server \
    && mv ./code-server-3.7.2-linux-amd64/* /usr/local/lib/code-server \
    && ln -s /usr/local/lib/code-server/code-server /usr/local/bin/code-server

RUN useradd -G sudo -ms /bin/bash zhixuan && \
    echo 'zhixuan:zhixuan' | chpasswd
RUN usermod -G sudo zhixuan

RUN adduser --gecos '' --disabled-password coder && \
    echo "coder ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers.d/nopasswd

ENV PORT=8080
EXPOSE 8080
USER coder
WORKDIR /home/coder

CMD code-server --host 0.0.0.0 --port $PORT .