awk 'BEGIN {FS="|"; OFS=" - "; sum=0; print "col1 - col2 - col3";} {print $1, $2, $3; sum+=$3;} END {print sum;}' data.txt
