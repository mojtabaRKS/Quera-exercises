groupadd shared
useradd -m -p "" -G shared user1
useradd -m -p "" -G shared user2

mkdir /shared_files
touch /shared_files/shared_file

chown user1:shared /shared_files
chown user1:shared /shared_files/shared_file

chmod 660 /shared_files/shared_file

userdel -r user1
userdel -r user2

groupdel shared

rm -r /shared_files



