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
    
elif [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"

    #ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    ip route add fc00:5::/64 via fd00:0:172:16:13::11 dev ${ETH}
    
    #ip addr add fc00:1::/64 dev lo
    #ip route add fc00:1:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute2" -o ${MODE} = "com2" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    #ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    ip route add fc00:5::/64 via fd00:0:172:16:13::11 dev ${ETH}

    #ip addr add fc00:2::/64 dev lo
    #ip route add fc00:2:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute3" -o ${MODE} = "com3" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    #ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    ip route add fc00:5::/64 via fd00:0:172:16:13::11 dev ${ETH}

    #ip addr add fc00:3::/64 dev lo
    #ip route add fc00:3:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute4" -o ${MODE} = "com4" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    #ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    ip route add fc00:5::/64 via fd00:0:172:16:13::11 dev ${ETH}

    #ip addr add fc00:4::/64 dev lo
    #ip route add fc00:4:: encap seg6local action End dev ${ETH}

elif [ ${MODE} = "compute5" -o ${MODE} = "com5" ]; then
    ./enable_seg6.py | sudo sh
    echo "----------------------------------------"
    ip route add fc00:1::/64 via fd00:0:172:16:2::4 dev ${ETH}
    ip route add fc00:2::/64 via fd00:0:172:16:4::11 dev ${ETH}
    ip route add fc00:3::/64 via fd00:0:172:16:4::12 dev ${ETH}
    ip route add fc00:4::/64 via fd00:0:172:16:5::11 dev ${ETH}
    #ip route add fc00:5::/64 via fd00:0:172:16:13::11 dev ${ETH}

    #ip addr add fc00:4::/64 dev lo
    #ip route add fc00:4:: encap seg6local action End dev ${ETH}

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