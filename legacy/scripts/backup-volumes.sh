#!/bin/bash

# Script to controller all Docker Volumes

echo "Starting Docker Volume Backup"

VOLUMES=($(docker volume ls --format "{{.Name}}"))
BACKUP_NAME=$(date +"%Y%m%d%H%M%S")

if mountpoint -q $BACKUP_DIR; then
    echo "NAS is mounted"
else
    echo "NAS is not mounted"
    exit 1
fi


mkdir $BACKUP_DIR/$BACKUP_NAME
cd $BACKUP_DIR/$BACKUP_NAME

for VOLUME_NAME in "${VOLUMES[@]}"
do
    VOLUME_TYPE=$(docker volume inspect "$VOLUME_NAME" --format "{{.Options.type}}")

    if [ "$VOLUME_TYPE" = "cifs" ] || [ "$VOLUME_TYPE" = "nfs" ]; then
        echo "Skipping volume: ${VOLUME_NAME}"
    else
        echo "Backing up: ${VOLUME_NAME}"
        docker run --rm -v $VOLUME_NAME:/data -v $BACKUP_VOLUME:/dest alpine:latest ash -c "tar cf /dest/$BACKUP_NAME/$VOLUME_NAME.backup.tar /data/."
    fi
done

cd $BACKUP_DIR

tar czf $BACKUP_NAME.tar.gz $BACKUP_NAME

rm -r $BACKUP_NAME

echo "Backup procedure was completed"