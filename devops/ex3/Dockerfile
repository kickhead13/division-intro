FROM alpine:latest

COPY main.c main.c
# COPY ptt.h ptt.h

# install gcc for compiling
RUN apk add build-base 

RUN gcc -o main main.c

ENTRYPOINT ["./main"]
