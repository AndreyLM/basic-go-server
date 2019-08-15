# main server settings
GCC=go
GCMD=run
GPATH=./cmd/server/server.go
RA512=4096

# db settings
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=4000
DB_DATABASE=basic
DB_USER=admin
DB_PASSWORD=admin123

# KEYS
KEYS_PATH=/home/andrew/Projects/go-modules/basic-server/internal/keys/
KEY_PRIVATE=${KEYS_PATH}app.rsa
KEY_PUBLIC=${KEYS_PATH}app.rsa.pub

DB_COMMAND_STRING= -db-driver=${DB_DRIVER} -db-host=${DB_HOST} -db-port=${DB_PORT} \
-db-database=${DB_DATABASE} -db-user=${DB_USER} -db-password=${DB_PASSWORD} \
-key-public=${KEY_PUBLIC} -key-private=${KEY_PRIVATE}

# jwt library doesn't use password so don't provide it or modify jwt package to use with password 
create_keys:
	-rm ${KEYS_PATH}*
	ssh-keygen -t rsa -b ${RA512} -f ${KEY_PRIVATE}
	openssl rsa -in ${KEY_PRIVATE} -pubout -outform PEM -out ${KEY_PUBLIC}

run:
	${GCC} ${GCMD} ${GPATH} ${DB_COMMAND_STRING}

show-flags:
	${GCC} ${GCMD} ${GPATH} -h