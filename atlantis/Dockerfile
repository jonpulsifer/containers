FROM ghcr.io/runatlantis/atlantis:v0.37.1@sha256:26043adad30edf8834da84d1e4ef4946e4c8b5222f287b0234d7edddfb136565
USER root
RUN echo https://downloads.1password.com/linux/alpinelinux/stable/ >> /etc/apk/repositories && \
    wget https://downloads.1password.com/linux/keys/alpinelinux/support@1password.com-61ddfc31.rsa.pub -P /etc/apk/keys && \
    apk add --no-cache 1password-cli && \
    op --version

# uid=100(atlantis) gid=1000(atlantis) groups=1000(atlantis)
USER 100:1000
