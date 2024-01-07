
1. Login into remote AWS server using the ssh command:`$ ssh ec2-user@ec2-ip-address-dns-name-here`
2. Apply pending updates using the [yum command](https://www.cyberciti.biz/faq/rhel-centos-fedora-linux-yum-command-howto/ "How to use yum command on CentOS/RHEL"):`$ sudo yum update`
3. Search for Docker package:`$ sudo yum search docker`
4. Get version information:`$ sudo yum info docker`  
5. Install docker, run:
	1. `$ sudo yum install docker`  
6. Add group membership for the default ec2-user so you can run all docker commands without using the sudo command:
	1. `sudo usermod -a -G docker ec2-user  
	2. `id ec2-user`
	3. `newgrp docker`
7. Need docker-compose too? Try any one of the following command:
    
    # 1. Get pip3 
    sudo yum install python3-pip
     
    # 2. Then run any one of the following
    sudo pip3 install docker-compose # with root access
     
    # OR #
     
    pip3 install --user docker-compose # without root access for security reasons
    
    OR
    
    wget https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) 
    sudo mv docker-compose-$(uname -s)-$(uname -m) /usr/local/bin/docker-compose
    sudo chmod -v +x /usr/local/bin/docker-compose
    
    [![Installing docker-compose on Amazon Linux 2 AMI](https://www.cyberciti.biz/media/new/faq/2021/09/Installing-docker-compose-on-Amazon-Linux-2-AMI-599x192.png)](https://www.cyberciti.biz/media/new/faq/2021/09/Installing-docker-compose-on-Amazon-Linux-2-AMI.png)
    
    How to install docker-compose in Amazon Linux (click to enlarge)
    
8. Enable docker service at AMI boot time:`$ sudo systemctl enable docker.service`
9. Start the Docker service:`$ sudo systemctl start docker.service`