FROM ubuntu:latest
LABEL authors="Pablo Pasquim"

ENTRYPOINT ["top", "-b"]