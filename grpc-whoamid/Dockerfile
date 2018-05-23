FROM alpine:latest
MAINTAINER John Belamaric <john@belamaric.com> @johnbelamaric

ADD grpc-whoamid /grpc-whoamid
ADD grpc-whoami /grpc-whoami

EXPOSE 8123
ENTRYPOINT ["/grpc-whoamid"]
