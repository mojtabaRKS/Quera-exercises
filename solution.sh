echo col1 - col2 - col3 && awk '{ gsub(/\|/, " - "); print ;  sum += $5 } END { print sum }' data.txt
