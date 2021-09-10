#!/bin/sh

echo
echo "###########################"
echo
echo "Starting SUT agents:"
echo

# check if exists
mkdir /var/log/tf-agent

for sutf in /data/testflinger-agent/sut/*; do
	if [ -f "$sutf" ]; then
		echo "$sutf"
		touch /var/log/tf-agent/${sutf%.*}
		testflinger-agent -c $sutf & > /var/log/tf-agent/${sutf%.*} 2>&1
		sleep .5
	fi
done

echo
echo "###########################"
echo