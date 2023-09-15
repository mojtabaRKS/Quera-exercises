
sudo useradd -m user1
sudo useradd -m user2
sudo groupadd shared 
sudo adduser user1 shared
sudo adduser user2 shared
sudo mkdir /root/shared_files
sudo touch /root/shared_files/shred_files
sudo chgrp -R shared /root/shared_files
sudo chown -R user1 /root/shared_files
sudo chmod 660 /root/shared_files/shared_files
sudo groupdel shared 
sudo userdel -r user1
sudo userdel -r user2
sudo rm -r /root/shared_files


