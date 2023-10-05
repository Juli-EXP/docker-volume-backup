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

ENV BACKUP_VOLUME=storage-location

ENV BACKUP_DIR=/backup

ENV TEMP_VOLUME=temp-location

ENV TEMP_DIR=/temp

ENV RETENTION=7

ENV IGNORED_VOLUMES=temp-location

ENV CRON_BACKUP=0 8 * * *

ENV CRON_PURGE=0 9 * * *

ENV CRON_TEMP=0 10 * * *


VOLUME /backup

VOLUME /temp

COPY scripts/*.sh /app/

RUN chmod +x /app/*.sh


COPY scripts/crontab /etc/cron.d/backup-volume

RUN chmod 644 /etc/cron.d/backup-volume


ENTRYPOINT [ "/app/entrypoint.sh" ]

CMD ["cron","-f", "-l", "2"]
