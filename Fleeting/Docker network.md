
----------------------------------
word:
- get handed
- I'll power through
- Ransomware Remediation
- something has back on that
- I gonna "take it out"
- once more
- I gonna do two things, first, .... second, ...
- leave it as is
- "specify" the nginx "at the very end"
- so far
- "bums off" the host
- essentially
- "even though" he is a container
- the "downside" is there's really no isolation
----------------------------------
## 1. default bridge
- copy the host's /etc/resolv.conf into container to make the same DNS as the host
- docker0 network interface act as a switch
- the containers on the bridge can use docker0 as the default gateway so that they can go to the internert

## 2. user defined bridge
- you can ping other containers which on the same bridge wich their DNS name
- To make the user defined bridge is a recommended way to deploy your container for the better isolation

## 3. host
- the container's network would use the host, you don't have to expose any ports
- the container would run as an regular application on the host

## 4. Mac-VLAN
- if there are too many mac address use the same port, you need to enable the Promisucous Mode.
- 


