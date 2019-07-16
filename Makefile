GCC=go
GCMD=run
GPATH=./cmd/server/main.go
RA512=4096

init:
	-rm ./internal/keys/*
	ssh-keygen -t rsa -b ${RA512} -f ./internal/keys/app.rsa
	openssl rsa -in ./internal/keys/app.rsa -pubout -outform PEM -out ./internal/keys/app.rsa.pub
run:
	${GCC} ${GCMD} ${GPATH}