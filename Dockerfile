FROM golang:alpine
WORKDIR /src
RUN apk add --update npm git
RUN go install github.com/go-bindata/go-bindata/...@latest
COPY ./webapp/package.json webapp/package.json
RUN cd ./webapp && \
    npm install --legacy-peer-deps
COPY . .
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN cd ./webapp && \
    npm run build
RUN cd ./migrations && \
    go-bindata  -pkg migrations .
RUN go-bindata  -pkg tmpl -o ./tmpl/bindata.go  ./tmpl/ && \
    go-bindata  -pkg webapp -o ./webapp/bindata.go  ./webapp/build/...    

RUN go build -o /opt/featmap/featmap && \
    chmod 775 /opt/featmap/featmap

ENTRYPOINT cd /opt/featmap && ./featmap
