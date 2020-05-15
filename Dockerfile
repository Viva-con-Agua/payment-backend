FROM golang

WORKDIR /go/src/payment-backend

ADD . /go/src/payment-backend
RUN ./install_packages.sh


CMD ["go", "run", "server.go"]
