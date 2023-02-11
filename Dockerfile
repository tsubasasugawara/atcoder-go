FROM golang:alpine

RUN apk update && \
	apk add vim git curl nodejs npm py3-pip bash

RUN pip3 install online-judge-tools && \
	npm install -g atcoder-cli

ARG USERNAME=user
ARG GROUPNAME=user
ARG UID=1000
ARG GID=1000
RUN addgroup -g $GID -S $GROUPNAME && \
    adduser -u $UID -S $USERNAME -G $GROUPNAME
USER $USERNAME
WORKDIR /home/$USERNAME

RUN acc config default-test-dirname-format test && \
	acc config default-task-choice all && \
	acc config default-template go

RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

COPY go ./.config/atcoder-cli-nodejs/go

COPY config/.bashrc ./.bashrc
COPY config/.vimrc ./.vimrc

COPY src ./go/src
WORKDIR ./go/src

CMD ["bash"]
