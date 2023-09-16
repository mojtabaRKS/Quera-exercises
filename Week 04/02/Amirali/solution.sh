groupadd shared
useradd -G shared -m -p "" user1
useradd -G shared -m -p "" user2
mkdir /shared_files
touch /shared_files/shared_file
chown -R user1:shared /shared_files
chmod 660 /shared_files/shared_file
deluser --remove-home user1
deluser --remove-home user2
groupdel shared
rm -Rf /home/user1
rm -Rf /home/user2
rm -Rf /shared_files