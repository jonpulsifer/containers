FROM gcr.io/trusted-builds/ubuntu-1804-base
LABEL maintainer=jonathan@pulsifer.ca \
      ca.lolwtf.os=linux \
      ca.lolwtf.distro=ubuntu \
      ca.lolwtf.codename=bionic \
      ca.lolwtf.release=18.04

ENV GO_VERSION=${GO_VERSION:-1.13.6}
ENV GO_CHECKSUM=${GO_CHECKSUM:-a1bc06deb070155c4f67c579f896a45eeda5a8fa54f35ba233304074c4abbbbd}

ARG LINUX_USER="${LINUX_USER:-jawn}"
ARG LINUX_UUID="${LINUX_UUID:-1337}"
ARG GITHUB_USER="${GITHUB_USER:-j0npulsifer}"
ARG AUTHORIZED_KEYS="/var/ssh/${LINUX_USER}/authorized_keys"

# user
RUN adduser \
  --disabled-password \
  --GECOS '' \
  --uid "${LINUX_UUID}" \
  "${LINUX_USER}"

WORKDIR /tmp/build

RUN rm /etc/apt/apt.conf.d/docker-no-languages
RUN apt-get update && apt-get -qy upgrade && \
    apt-get -qy install \
      apt-transport-https \
      ca-certificates \
      curl \
      lsb-release \
      software-properties-common \
      bash-completion \
      cowsay \
      dnsutils \
      figlet \
      fortune \
      gawk \
      gcc \
      git \
      inetutils-traceroute \
      iputils-ping \
      iputils-tracepath \
      jq \
      lolcat \
      make \
      man \
      netcat \
      nmap \
      openssh-client \
      openssh-server \
      python-dev \
      python-pip \
      rsync \
      shellcheck \
      sudo \
      telnet \
      tmux \
      unzip \
      vim \
      wget

# sudo and ssh setup
RUN echo "${LINUX_USER}     ALL=(ALL:ALL) NOPASSWD: ALL" >> /etc/sudoers && \
    mkdir -vp /var/ssh/"${LINUX_USER}" && \
    curl -sOJ https://github.com/"${GITHUB_USER}".keys && \
    mv -v "${GITHUB_USER}".keys "${AUTHORIZED_KEYS}" && \
    chown "${LINUX_USER}":"${LINUX_USER}" "${AUTHORIZED_KEYS}" && \
    chmod 644 "${AUTHORIZED_KEYS}"

# golang
ADD https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz .
RUN echo "${GO_CHECKSUM} *go${GO_VERSION}.linux-amd64.tar.gz" | sha256sum -c - && \
    tar -xzf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local/

# gcloud
ADD https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.tar.gz .
RUN tar -xzf google-cloud-sdk.tar.gz -C /usr/local/ && \
    /usr/local/google-cloud-sdk/install.sh --quiet --additional-components alpha beta

# set up ssh
COPY etc/ssh/sshd_config /etc/ssh/sshd_config
COPY etc/environment /etc/environment
RUN mkdir -p /run/sshd

USER ${LINUX_UUID}:${LINUX_UUID}
WORKDIR /home/${LINUX_USER}
