ifconfig |  grep -v 127 |grep -v  inet6 |grep inet   |awk {'print $2'}
