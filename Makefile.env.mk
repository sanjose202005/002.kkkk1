
CFGmakeRun:=Makefile.run.go01.mk

GoTOP:=\
	httpproxy2_main

httpproxy2_main := \
	httpproxy2_info01 	\
	httpproxy2_access 	\
	httpproxy2_https    \
	httpproxy2_httP    	\
	\
	httpproxy2_main

# https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c
sslConfig:=/etc/ssl/openssl.cnf
ssl_conf:
	[ -d ssl ] || mkdir ssl
	[ -f ssl/myOpenssl.conf ] \
		&& printf '\nAlread found : ssl/myOpenssl.conf , skip. \n\n' \
		|| ( \
		cat $(sslConfig)                              > ssl/myOpenssl.conf && \
		printf '[SAN]\nsubjectAltName=DNS:localhost' >> ssl/myOpenssl.conf \
		)
	[ -f ssl/myServer.key ] \
		&& printf '\nAlread found : ssl/myServer.key , skip. \n\n' \
		|| \
		openssl req \
		-newkey rsa:2048 \
		-x509 \
		-nodes \
		-keyout ssl/myServer.key \
		-new \
		-out ssl/myServer.pem \
		-subj /CN=localhost \
		-reqexts SAN \
		-extensions SAN \
		-config                 ssl/myOpenssl.conf    \
		-sha256 \
		-days 3650
