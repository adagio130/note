
1. Login into remote AWS server using the ssh command:`$ ssh ec2-user@ec2-ip-address-dns-name-here`
2. Apply pending updates using the yum command:
	1. `$ sudo yum update`
3. Search for Docker package:
	1. `$ sudo yum search docker`
4. Get version information:
	1. `$ sudo yum info docker`  
5. Install docker, run:
	1. `$ sudo yum install docker`  
6. Add group membership for the default ec2-user so you can run all docker commands without using the sudo command:
	1. `sudo usermod -a -G docker ec2-user  
	2. `id ec2-user`
	3. `newgrp docker`
7. Need docker-compose too? Try any one of the following command:
	1. Get pip3 
		1. `sudo yum install python3-pip`
	2. Then run any one of the following
		1. `sudo pip3 install docker-compose # with root access`
		# or
		2. `pip3 install --user docker-compose # without root access for security reasons`
8. Start the Docker service:
	1. `sudo systemctl start docker.service`
9. Enable docker service at AMI boot time:`
	1. `sudo systemctl enable docker.service`