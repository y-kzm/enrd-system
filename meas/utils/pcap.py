#!/bin/env python3
from ast import keyword
from genericpath import samestat
from dbus import Interface
import yaml
import subprocess
import sys

args = sys.argv
argc = len(args)
if argc != 3:
    print('Argment error')
    print('./pcap.py [yaml file] [pcap path]')
    quit()
file_name = args[1]
pcap_dir = args[2]
print("Start!")
 
with open(file_name) as file:
    yml = yaml.safe_load(file)
    nodes = yml['nodes'] 
    name = [d.get('name') for d in nodes]   # ['R1', 'R2', 'R3', ...]
for i in name:
    try:
        subprocess.check_output(["docker", "exec", i, "pkill", "tcpdump"])
        cmd_rslt = subprocess.check_output(["docker", "exec", i, "ls", "/tmp"])
        cmd_rslt = cmd_rslt.decode().strip().split('\n')  
        sampling = []; key = ".pcap"
        try:
            for j in cmd_rslt:
                if key in j:
                    sampling.append(j)
            for k in sampling:
                subprocess.check_output(["docker", "cp", i + ":/tmp/" + k, pcap_dir + k])
        except:
            print("Error")
    except:
        print(i + ': tcpdump does not exist.')

print("Finish!")
quit()