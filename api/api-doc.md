# Backup information
GET     /api/backup/                     Returns array of all volumes with corresponding backup information  
GET     /api/backup/{volume-name}        Returns backup information of single volume  

# Backup creation
POST    /api/backup/create                Creates backup of all volumes, nfs and cifs are disabled by default  
POST    /api/backup/create/{volume-name}  Creates backup of a single volume  

# Backup purge
POST    /api/backup/purge                       Deletes all volume backups older than x days
POST    /api/backup/purge/{volume-name}         Deletes volume backups of a singel volume older than x days

# Backup deletion
DELETE    /api/backup/delete/{backup name / id}   Deletes a specific volume backup

# Backup configuration
GET     /api/config

PUT /api/config/x