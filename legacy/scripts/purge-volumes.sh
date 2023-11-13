#!/bin/bash

# Script to delete volume backups older than 2 weeks

echo "Start of Docker Volume purge"

THRESHOLD=$(date -d "$RETENTION days ago" +"%Y%m%d%H%M%S")

find ${BACKUP_DIR} -maxdepth 1 -type f -print0  | while IFS= read -d '' -r file
do
    #Check if filename matches controller filename structure
    if [[ "$(basename "$file")" =~ ^.*.tar\.gz$ ]]
    then
        # Check if file is old enough to delete it
        if [[ "$(basename "$file" .tar.gz)" -le "$THRESHOLD" ]]
        then
            echo "Deleting $(basename "$file")"
            rm -r $file
        fi
    fi
done

echo "Old Docker Volumes were purged"