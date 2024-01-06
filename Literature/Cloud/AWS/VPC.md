![[截圖 2024-01-06 上午9.48.41.png]]
1. VPC
	- manage the internet traffic
	- subnets
		- like the real subnets, put resources inside to manage them.
		- public subnet
		- private subnet
			- close ne
	- Routing table
		- each subnet would has a routing table
1. DHCP options
2. EIPs
3. ENIs
4. VPC Peering
   - connects one vpc to another
   - can't transitive
   - creating vpc peers
     - owner role required
     - IP CIDR blocks in each VPC must not overlap
     - VPC do not use the same IP CIDR
     - each VPC needs a defined route to the other VPC
     - security group rules
5. Security group (like firewall)
     - assigned to an instance in a VPC
     - applied to instances not to subnets
     - Ingress (entrance、incoming)
     - Egress (exit、outgoing)
     - only Allow
6. Network Access Control Lists(NACLs)
     - Applied on subnets
     - Both Allow/deny
     - Lowest numbered rules first
     - First match applies
     - like routers ACL
7. NAT
     - translates between Private IP address and Public Address 
8. VPG
     - connects local networks to the vpc
     - VPG is the VPN concentrator