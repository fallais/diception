FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"

ADD diception /usr/bin

EXPOSE     8000
CMD        [ "/usr/bin/diception" ]
