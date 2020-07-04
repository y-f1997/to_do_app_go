FROM golang:1.13

USER root

RUN apt update \
&& apt install -y nodejs npm \
&& npm install n -g \
&& n stable \
&& apt purge -y nodejs npm 

COPY . /go/
WORKDIR /go/
RUN go get github.com/coopernurse/gorp \
&& go get github.com/dgrijalva/jwt-go \
&& go get github.com/go-gorp/gorp \
&& go get github.com/go-ozzo/ozzo-validation \
&& go get github.com/lib/pq \
&& go get -u golang.org/x/crypto/...

WORKDIR /go/src/go_app/views
RUN npm install && npm run build
WORKDIR /go/src/go_app/

CMD ["go","run", "main.go"]
