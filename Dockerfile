FROM alpine:latest
COPY ./myapp /build/bin/myapp

EXPOSE 8080

CMD "/build/bin/myapp"