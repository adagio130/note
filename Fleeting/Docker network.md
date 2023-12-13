

## the default bridge
- copy the host's /etc/resolv.conf into container to make the same DNS as the host
- docker0 network interface act as a switch
- the containers on the bridge can use docker0 as the default gateway so that they can go to the internert

## the user defined bridge
