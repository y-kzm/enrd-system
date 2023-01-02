#!/bin/bash
#MODE=${1}
ETH=${1}
MODE=`hostname`

if [ $# != 1 ]; then
        echo "Usage: ./${PROGRAM} [Interface]"
        exit 0;
fi

if [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
    echo "----------------------------------------"
    ip addr del fc00:1::/64 dev lo
    ip route del fc00:1:: encap seg6local action End dev ${ETH} 

elif [ ${MODE} = "compute2" -o ${MODE} = "com2" ]; then
    echo "----------------------------------------"
    ip addr del fc00:2::/64 dev lo
    ip route del fc00:2:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute3" -o ${MODE} = "com3" ]; then
    echo "----------------------------------------"
    ip addr del fc00:3::/64 dev lo
    ip route del fc00:3:: encap seg6local action End dev ${ETH} 

elif [ ${MODE} = "compute4" -o ${MODE} = "com4" ]; then
    echo "----------------------------------------"
    ip addr del fc00:4::/64 dev lo
    ip route del fc00:4:: encap seg6local action End dev ${ETH} 

else
    echo "Not supported"
    exit 0;
fi

if [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
    echo "----------------------------------------"
    ip -6 rule del from fd00:0:172:16:ffff::1/64
    ip -6 rule del from fd00:0:172:16:ffff::2/64
    ip -6 rule del from fd00:0:172:16:ffff::3/64

    ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:2:: dev ${ETH} table 100 
    ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:3:: dev ${ETH} table 101 
    ip -6 route del fc00:4::/64 encap seg6 mode encap segs fc00:2::,fc00:3:: dev ${ETH} table 102

    ip addr del fd00:0:172:16:ffff::1/64 dev ${ETH}
    ip addr del fd00:0:172:16:ffff::2/64 dev ${ETH}
    ip addr del fd00:0:172:16:ffff::3/64 dev ${ETH}

else
    echo "Nothing..."
    exit 0;
fi

echo "-------------------------"
ip -6 rule show
echo "-------------------------"
ip -6 route show table all
echo "-------------------------"
ip addr show lo
ip addr show ${ETH}
