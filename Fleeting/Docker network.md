
----------------------------------
word:
- get handed
- I'll power through
- Ransomware Remediation
- something has back on that

----------------------------------
## 1. the default bridge
- copy the host's /etc/resolv.conf into container to make the same DNS as the host
- docker0 network interface act as a switch
- the containers on the bridge can use docker0 as the default gateway so that they can go to the internert

## 2. the user defined bridge
- you can ping other containers which on the same bridge wich their DNS name
- For the better isolmake the user defined bridge is a recommended way to deploy your container
