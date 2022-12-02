FROM yykzm/ubuntu:20.04

COPY ./igi-ptr-2.1-v6 /v6
WORKDIR /v6
RUN make \ 
 && chmod +x auto.sh

CMD ["/bin/bash"]

#FROM slankdev/ubuntu:18.04
#
#RUN apt update && apt install -y iperf3
#COPY ./igi-ptr-2.1-v6 /v6
#WORKDIR /v6
#RUN make \ 
# && chmod +x auto.sh
#
#CMD ["/bin/bash"]