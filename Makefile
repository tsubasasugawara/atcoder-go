build:
	docker build -t alpine:golang .

run:
	docker run -it -v $(PWD)/src:/go/src \
	-v $(PWD)/go:/home/user/.config/atcoder-cli-nodejs/go \
   	-w /go/src --name atcoder-go alpine:golang

start:
	clear
	docker cp config/.vimrc atcoder-go:/home/user/.vimrc
	docker cp config/.bashrc atcoder-go:/home/user/.bashrc
	docker start -ia atcoder-go
