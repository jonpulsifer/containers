FROM gcr.io/trusted-builds/ubuntu-1804-base

ENV PACKAGES build-essential libncurses-dev bison flex libssl-dev libelf-dev

RUN apt-get -qqy update && apt-get -qqy upgrade && \
    apt-get -qqy install ${PACKAGES}

# google cloud build
WORKDIR /workspace
