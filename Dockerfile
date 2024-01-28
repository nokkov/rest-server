FROM golang:1.21

WORKDIR /rest-server

COPY . .

EXPOSE 8084

RUN go get -d . \
    && go install

ENV CONFIG_PATH config/local.yaml

CMD ["go", "run", "main.go"]