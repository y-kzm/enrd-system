#!/bin/bash
#MODE=${1}
ETH=${1}
MODE=`hostname`

if [ $# != 1 ]; then
        echo "Usage: ./${PROGRAM} [Interface]"
        exit 0;
fi

if [ ${MODE} = "controller" -o ${MODE} = "con" ]; then
    echo "----------------------------------------"
    echo "Nothing..."
    # ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev ${ETH}
    # ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev ${ETH}
    # ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev ${ETH}
    # ip route add fc00:4::/64 via fd00:0:172:16::5:11 dev ${ETH}
elif [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"

    #ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    
    #ip addr add fc00:1::/64 dev lo
    #ip route add fc00:1:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute2" -o ${MODE} = "com2" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    #ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}

    #ip addr add fc00:2::/64 dev lo
    #ip route add fc00:2:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute3" -o ${MODE} = "com3" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    #ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}

    #ip addr add fc00:3::/64 dev lo
    #ip route add fc00:3:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute4" -o ${MODE} = "com4" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    #ip route add fc00:4::/64 via fd00:0:172:16::5:11 dev ${ETH}

    #ip addr add fc00:4::/64 dev lo
    #ip route add fc00:4:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "config" ]; then
    echo "----------------------------------------"
    ip addr add fd00:0:172:16:ffff::1/64 dev ${ETH}
    ip addr add fd00:0:172:16:ffff::2/64 dev ${ETH}
    ip addr add fd00:0:172:16:ffff::3/64 dev ${ETH}

    ip -6 rule add from fd00:0:172:16:ffff::1 table 100
    ip -6 rule add from fd00:0:172:16:ffff::2 table 101
    ip -6 rule add from fd00:0:172:16:ffff::3 table 102

    ip route add fc00:4::/64 encap seg6 mode encap segs fc00:2:: dev ${ETH} table 100
    ip route add fc00:4::/64 encap seg6 mode encap segs fc00:3:: dev ${ETH} table 101
    ip route add fc00:4::/64 encap seg6 mode encap segs fc00:2::,fc00:3:: dev ${ETH} table 102

else
    echo "Not supported"
    exit 0;
fi

echo "-------------------------"
ip -6 rule show
echo "-------------------------"
ip -6 route show table all
echo "-------------------------"
ip addr show lo
ip addr show ${ETH}