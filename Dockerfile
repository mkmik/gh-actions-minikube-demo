FROM golang:1.15 as builder

WORKDIR /src

COPY . .

RUN go build -tags netgo -ldflags '-w'

FROM gcr.io/distroless/static:nonroot@sha256:75c99ae4ddc137571997e6e95f78c30dca226118326a0755ad19ebc4955abe3d
MAINTAINER Marko Mikulicic <mmikulicic@gmail.com>
COPY --from=builder /src/gh-actions-minikube-demo /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/gh-actions-minikube-demo"]
