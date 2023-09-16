groupadd shared 
useradd -G shared -m  user1
useradd -G shared -m  user2
mkdir /shared_files
touch /shared_files/shared_file
chgrp -R shared /shared_files
chown -R user1 /shared_files
chmod 660 /shared_files/shared_files
groupdel shared 
userdel -r user1
userdel -r user2
rm -r /shared_files

