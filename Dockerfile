FROM scratch
EXPOSE 8080
ENTRYPOINT ["/hello-go"]
COPY ./bin/ /