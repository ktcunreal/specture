#!/usr/bin/env bash

# Check privilege
if [[ ! `whoami` == "root" ]]; then
	echo "SU privilege is required";
	exit 1;
fi

# Ip check
if [ -z $1 ]; then
	echo "Empty ip is not allowed";
	exit 1;
fi

# Check Path
if [[ ! -f $2 ]]; then
	echo "Invalid whitelist path!"
	exit 1;
fi

# Check if ip is duplicated
if [[ -n "`grep -w $1 $2`" ]]; then
	echo "IP seems already exist, please check /etc/haproxy/whitelist for detail";
	exit 1;
fi


# Add ip to whitelist
echo "$1/32" >> $2

# Purge IP in other whitelist
for ln in `grep "acl whitelist" /etc/haproxy/haproxy.cfg  | awk '{print $NF}' | grep -vw $2`; do 
	sed -i "/$1/d" $ln
done

# Print result
echo "$1/32 added"

systemctl reload haproxy
