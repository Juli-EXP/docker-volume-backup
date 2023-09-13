FROM debian:bookworm-slim

ARG DOCKER_CLIENT="docker-24.0.6.tgz"

RUN apt update \
&& apt upgrade -y \
&& apt install curl cron -y 

# Install the docker client binaries
RUN cd /tmp/ \
&& curl -sSL -O https://download.docker.com/linux/static/stable/x86_64/${DOCKER_CLIENT} \
&& tar zxf ${DOCKER_CLIENT} \
&& mkdir -p /usr/local/bin \
&& mv ./docker/docker /usr/local/bin \
&& chmod +x /usr/local/bin/docker \
&& rm -rf /tmp/*

ENV DOCKER_HOST=unix:///var/run/docker.sock

ENV BACKUP_VOLUME=backup-data

ENV BACKUP_DIR=/backup

ENV RETENTION=7

VOLUME /backup

COPY scripts/*.sh /app/

COPY scripts/crontab /etc/crontab

RUN chmod +x /app/*.sh

ENTRYPOINT [ "/app/entrypoint.sh" ]

CMD ["cron","-f", "-l", "2"]