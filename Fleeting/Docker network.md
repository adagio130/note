
default:
	bridget
		- copy the host's /etc/resolv.conf to container into resolve the DNS
		- docker 0 network interface act as a switch
		- the containers on the bridge can use docker0 as the default gateway so that they can go to the internert
		- 