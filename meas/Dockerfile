FROM slankdev/ubuntu:18.04

RUN apt update && apt upgrade -y
RUN apt install -y iperf3 net-tools
COPY ./igi-ptr-2.1-origin /origin
COPY ./igi-ptr-2.1-v6 /v6


