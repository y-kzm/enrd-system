#!/bin/bash
MODE=${1}
if [ $# != 1 ]; then
        echo "Usage: ./${PROGRAM} [mode]"
        echo "mode: controller, comptue1, ..., compute4"
        exit 0;
fi

elif [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
        ip -6 rule del from fd00:0:172:16::ffff:1/64
        ip -6 rule del from fd00:0:172:16::ffff:2/64
        ip -6 rule del from fd00:0:172:16::ffff:3/64

        ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:2:: dev enp6s0 table 100 metric 1024 pref medium
        ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:3:: dev enp6s0 table 101 metric 1024 pref medium
        ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:3::,fc00:2:: dev enp6s0 table 102 metric 1024 pref medium

        ip addr del fd00:0:172:16::ffff:1/64 dev enp6s0
        ip addr del fd00:0:172:16::ffff:2/64 dev enp6s0
        ip addr del fd00:0:172:16::ffff:3/64 dev enp6s0

        ip addr del fc00:1::/64 dev lo
        ip route del fc00:1::/64

elif [ ${MODE} = "compute2" -o ${MODE} = "com2" ]; then
        ip addr del fc00:2::/64 dev lo
        ip route del fc00:2::/64

elif [ ${MODE} = "compute3" -o ${MODE} = "com3" ]; then
        ip addr del fc00:3::/64 dev lo
        ip route del fc00:3::/64

elif [ ${MODE} = "compute4" -o ${MODE} = "com4" ]; then
        ip addr del fc00:4::/64 dev lo
        ip route del fc00:4::/64
else
    echo "Not supported"
    exit 0;
fi

echo "-------------------------"
ip -6 route show table all
echo "-------------------------"
ip -6 route show
echo "-------------------------"
ip addr show lo
ip addr show enp6s0
