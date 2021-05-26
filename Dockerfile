FROM scratch

COPY ./bin/faas /faas

CMD ["/faas"]
