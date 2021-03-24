FROM ubuntu:focal

ENV TZ Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update && apt-get install -y \
  zsh \
  build-essential \
  dnsutils \
  netcat \
  net-tools \
  nmap \
  curl \
  traceroute \
  tcpdump \
  htop \
  iftop \
  iotop \
  iputils-ping \
  mtr-tiny \
  vim \
  git \
  && rm -rf /var/lib/apt/lists/*
RUN git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime \
  && sh ~/.vim_runtime/install_awesome_vimrc.sh
RUN mkdir /playground

COPY playground /usr/local/bin

WORKDIR /playground
CMD     ["playground"]
