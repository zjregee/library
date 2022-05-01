FROM alpine

COPY ./app /root/lib

EXPOSE 8080

ENTRYPOINT ["/root/lib"]
