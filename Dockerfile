FROM golang:1.9-stretch
MAINTAINER Dan Itsara <dan@glazziq.com>
RUN apt-get update -qq && \
  apt-get install -y less && \
  curl https://glide.sh/get | sh && \
  go get -u github.com/onsi/ginkgo/ginkgo
