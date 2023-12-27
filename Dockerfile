FROM golang

WORKDIR /url-shortener

ENV CONFIG_PATH = $WORKDIR/config/local.yaml

COPY . . 
# TODO: VOLUME here

RUN go get -d ./...

EXPOSE 8084

CMD ["go", "run", "main.go"]