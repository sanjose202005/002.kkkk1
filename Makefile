srcS:=valuga https_baidu httpproxy http_fileserver


define helpText

https://www.reddit.com/r/golang/comments/8m4xrh/do_linux_golang_binaries_depend_on_libc/
	CGO_ENABLED=0 go build -o bin.bin src.go
  User/group resolution, and the non-DNS aspects for host resolution, 
are code that only exists in the libc, and the way glibc implements 
those is by dynamically loading libraries based on config files

b1: $(src1) 
b2: $(src2) 
b3: $(src3) 

d1 d2 d3 --> run in docker

endef

export helpText

all:
	@echo "$${helpText}"

m:
	vim Makefile
gs:
	git status
gc:
	git commit -a

ga:
	git add .

r1 v1 b1 d1: src:=$(src1)
r2 v2 b2 d2: src:=$(src2)
r3 v3 b3 d3: src:=$(src3)

b1 b2 b3: 
	@echo ; echo === $@
	rm -f       bin.$(src).bin 
	CGO_ENABLED=0 go build -o bin.$(src).bin $(src).go
	@strip                    bin.$(src).bin 
	@ls -l                    bin.$(src).bin 
	@file                     bin.$(src).bin 
	@echo

v1 v2 v3:
	vim $(src).go

r1 r2 r3:
	./bin.$(src).bin

all1 : v1 b1 r1
all2 : v2 b2 r2
all3 : v3 b3 r3

# https://docs.docker.com/engine/reference/commandline/run/
d1 d2 d3 :
	@echo docker exec -it  XXX sh
	docker run \
		--read-only \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		-i \
		-t \
		-d \
		--rm \
		--memory 100m \
		--network host \
		busybox \
		nice -n 19 ./bin.$(src).bin

define if_home_end_key_do_NOT_work_in_docker
https://superuser.com/questions/94436/how-to-configure-putty-so-that-home-end-pgup-pgdn-work-properly-in-bash
~/.inputrc

set meta-flag on
set input-meta on
set convert-meta off
set output-meta on
"\e[1~": beginning-of-line     # Home key
"\e[4~": end-of-line           # End key
"\e[5~": beginning-of-history  # PageUp key
"\e[6~": end-of-history        # PageDown key
"\e[3~": delete-char           # Delete key
"\e[2~": quoted-insert         # Insert key
"\eOD": backward-word          # Ctrl + Left Arrow key
"\eOC": forward-word           # Ctrl + Right Arrow key

endef
