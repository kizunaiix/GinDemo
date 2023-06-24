FROM alpine:3.14

RUN mkdir "/app"
WORKDIR "/app"

COPY GinDemo /app/app
COPY resources /app/resources

ENTRYPOINT [ "./app" ]
