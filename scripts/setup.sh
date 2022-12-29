#!/bin/bash
MODE=${1}
if [ $# != 1 ]; then
        echo "Usage: ./${PROGRAM} [mode]"
        echo "mode: controller, comptue1, ..., compute4"
        exit 0;
fi

if [ ${MODE} = "controller" -o ${MODE} = "con" ]; then
    ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev enp6s0
    ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev enp6s0
    ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev enp6s0
    ip route add fc00:4::/64 via fd00:0:172:16::11:5 dev enp6s0
elif [ ${MODE} = "compute1" -o ${MODE} = "com1" ]; then
    #ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev enp6s0
    ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev enp6s0
    ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev enp6s0
    ip route add fc00:4::/64 via fd00:0:172:16::11:5 dev enp6s0
    
    ip addr add fc00:1::/64 dev lo
    ip route add fc00:1::/128 encap seg6local action End dev enp6s0
elif [ ${MODE} = "compute2" -o ${MODE} = "com2" ]; then
    ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev enp6s0
    #ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev enp6s0
    ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev enp6s0
    ip route add fc00:4::/64 via fd00:0:172:16::11:5 dev enp6s0

    ip addr add fc00:2::/64 dev lo
    ip route add fc00:2::/128 encap seg6local action End dev enp6s0
elif [ ${MODE} = "compute3" -o ${MODE} = "com3" ]; then
    ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev enp6s0
    ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev enp6s0
    #ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev enp6s0
    ip route add fc00:4::/64 via fd00:0:172:16::11:5 dev enp6s0

    ip addr add fc00:3::/64 dev lo
    ip route add fc00:3::/128 encap seg6local action End dev enp6s0
elif [ ${MODE} = "compute4" -o ${MODE} = "com4" ]; then
    ip route add fc00:1::/64 via fd00:0:172:16::2:4 dev enp6s0
    ip route add fc00:2::/64 via fd00:0:172:16::4:11 dev enp6s0
    ip route add fc00:3::/64 via fd00:0:172:16::4:12 dev enp6s0
    #ip route add fc00:4::/64 via fd00:0:172:16::11:5 dev enp6s0

    ip addr add fc00:4::/64 dev lo
    ip route add fc00:4::/128 encap seg6local action End dev enp6s0
else
    echo "Not supported"
    exit 0;
fi

echo "------------"
ip addr show lo
ip addr show enp6s0
echo "------------"
ip -6 route show
