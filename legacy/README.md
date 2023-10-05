# docker-volume-backup

A Docker Container which backs up the content of all local Docker Volumes


## Information

This docker container executes the following scripts through cron:
- backup-volumes.sh  
    - Backs up all local docker volumes and puts them in a tarball named after the current timestamp
        ```
        |---202309170800.tar.gz
        |   |---volume-1.tar
        |   |---volume-2.tar
        |   |   ...
        ```
        
    - This will ignore nfs and cifs mounted volumes An option to enable backups of these volume types might be possible in the future 

- purge-volumes.sh
    - Deletes all backups older than the retention time (in days)

- check-temp.sh
    - If the main location for the backups, usually some kind of network storage, is not available, the backups are stored in a temp folder. 
    - The script checks if the main backup location is available again and it moves the backups there


## Usage

### docker-compose.yml

```yaml
version: "3.9"

services:
  docker-volume-backup:
    image: docker-volume-backup:latest
    networks:
      network:  # Replace with your network
      docker-socket_network:
    environment:
      - TZ=Europe/Rome  # Adjust to your timezone
      - DOCKER_HOST=tcp://docker-socket-proxy:2375
      - BACKUP_VOLUME=docker-volume-backup_storage-location # Use the full name of the docker volume which gets displayed by running docker volume ls
      - BACKUP_DIR=/backup  # Optional: change the location of the backup folder inside the container. Needs to match the volume mounting point
      - TEMP_VOLUME=docker-volume-backup_temp-location # Use the full name of the docker volume which gets displayed by running docker volume ls
      - TEMP_DIR=/temp  # Optional: change the location of the temp folder inside the container. Needs to match the volume mounting point. Folder is used if /backup is not available
      - RETENTION=7 # Optional: change the number of days of backups to be kept
      - IGNORED_VOLUMES=docker-volume-backup_temp-location  # Comma seperated list of volumes wich are ignored by the backup script
      - CRON_BACKUP=0 8 * * *,@reboot # Optional: comma seperated list of cron schedules for the backup script
      - CRON_PRUGE=0 9 * * *  # Optional: comma seperated list of cron schedules for the purge script
      - CRON_TEMP=0 10 * * *  # Optional: comma seperated list of cron schedules for the temp folder script
    volumes:
      - storage-location:/backup  # Replace with your volume
      - temp-location:/temp # Replace with your volume
    depends_on:
      - docker-socket-proxy
    restart: always

  docker-socket-proxy:
    image: ghcr.io/tecnativa/docker-socket-proxy:0.1.1
    networks: 
      docker-socket_network:
    environment:
      - VOLUMES=1 # Listing volumes
      - IMAGES=1  # Pulling images
      - CONTAINERS=1  # Creating containers if POST is set to 1
      - POST=1
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro 
    restart: always


volumes:
  storage-location: # Replace with your volume configuration
    driver: local
    driver_opts:
      type: cifs
      device: "//my/remote/storage"
      o: username=${CIFS_USERNAME},password=${CIFS_PASSWORD},addr=my-address,vers=3.0
  temp-location:  # Replace with your volume configuration

networks:
  network:  # Replace with your network
    external: true  # Optional: only if network was created outside of the docker-compose file
  docker-socket_network:  # Communication between the container and the docker socket proxy
    internal: true
```

## Parameters

### Environment Variables

| Name | Default Value | Function |
| --- | --- | --- |
| DOCKER_HOST | unix:///var/run/docker.sock | Location of the docker socket can also be a network address, e.g.: tcp://docker-socket-proxy:2375 |
| BACKUP_VOLUME | backup-data | The full name of the backup volume which gets displayed by running docker volume ls |
| BACKUP_DIR | /backup | Location of the backup directory inside the container |
| TEMP_VOLUME | temp-location | The full name of the temp volume which gets displayed by running docker volume ls |
| TEMP_DIR | /temp | Location of the backup directory inside the container |
| RETENTION | 7 | Number of days how long the backups will be kept |
| IGNORED_VOLUMES | docker-volume-backup_temp-location | Comma seperated list of volumes wich are ignored by the backup script
| CRON_BACKUP | 0 8 * * * | Comma seperated list of cron schedules for the backup script |
| CRON_PURGE | 0 9 * * * | Comma seperated list of cron schedules for the purge script |
| CRON_TEMP | 0 10 * * * | Comma seperated list of cron schedules for the temp folder script |


### Volume mappings


| Volume | Function |
| --- | --- |
| /backup | Primary location for the backup storage. It is recomended to use a nfs or cifs volume |
| /temp | Secondary location for the backup storage, in case the primary is down. It is recomended to use a local volume |

