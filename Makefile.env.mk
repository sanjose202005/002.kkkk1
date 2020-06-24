
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
	openssl req \
		-newkey rsa:2048 \
		-x509 \
		-nodes \
		-keyout server.key \
		-new \
		-out server.pem \
		-subj /CN=localhost \
		-reqexts SAN \
		-extensions SAN \
		-config <(cat $(sslConfig) \
		<(printf '[SAN]\nsubjectAltName=DNS:localhost')) \
		-sha256 \
		-days 3650
