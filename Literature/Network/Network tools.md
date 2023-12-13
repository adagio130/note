
# Check the port status

## host file
`grep -w '80/tcp' /etc/services`
`grep -w '443/tcp' /etc/services`
`grep -E -w '22/(tcp|udp)' /etc/services`
## netstat
`sudo netstat -tulpn | grep LISTEN`
`cat /etc/services`


# tcpdump

