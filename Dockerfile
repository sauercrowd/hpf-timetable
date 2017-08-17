FROM golang:1.8
MAINTAINER sauercrowd <jonadev95@posteo.org>

ADD . /go/src/github.com/sauercrowd/hpf-timetable
RUN go install github.com/sauercrowd/hpf-timetable
WORKDIR /go/src/github.com/sauercrowd/hpf-timetable
EXPOSE 8080