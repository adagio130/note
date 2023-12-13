
----------------------------------
word:
- get handed
- I'll power through
- Ransomware Remediation
- something has back on that
- I gonna take it out
- once more
- I gonna do two things, first, .... second, ...
- leave it as is
- specify the nginx at the very end
- so far
- bums off the host
----------------------------------
## 1. the default bridge
- copy the host's /etc/resolv.conf into container to make the same DNS as the host
- docker0 network interface act as a switch
- the containers on the bridge can use docker0 as the default gateway so that they can go to the internert

## 2. the user defined bridge
- you can ping other containers which on the same bridge wich their DNS name
- To make the user defined bridge is a recommended way to deploy your container for the better isolation

## 3. the host
- you don't have to expose any ports
