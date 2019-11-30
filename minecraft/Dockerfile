FROM adoptopenjdk/openjdk11:alpine

ARG MINECRAFT_VERSION=${MINECRAFT_VERSION:-1.14.4}
ARG MINECRAFT_SERVER_CHECKSUM=${CHECKSUM:-3dc3d84a581f14691199cf6831b71ed1296a9fdf}
ARG NAME=${NAME:-minecraft}
ARG APPDIR=${APPDIR:-/var/minecraft}
ARG UID=${UID:-1000}
ARG GID=${GID:-1000}

RUN addgroup -g ${GID} ${NAME} && \
    adduser -D -u ${UID} -G ${NAME} ${NAME}

RUN mkdir -p ${APPDIR} && chmod 700 ${APPDIR}

ADD https://launcher.mojang.com/v1/objects/${MINECRAFT_SERVER_CHECKSUM}/server.jar /usr/share/java/minecraft.jar
RUN chmod 644 /usr/share/java/minecraft.jar && chown -R ${UID}:${GID} ${APPDIR}

WORKDIR ${APPDIR}
USER ${UID}:${GID}
RUN echo "eula=true" > eula.txt

ENTRYPOINT ["java", "-jar", "/usr/share/java/minecraft.jar", "nogui"]
