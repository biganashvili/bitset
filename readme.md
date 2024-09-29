# Counting the number of unique IP addresses in a very large file

In theory, there can be 2^32 unique IP addresses.

1) Using map[uint32]bool to store the occurred IP addresses would be inefficient.
since uint32 uses 4 bytes and bool uses 1 byte in go. 2^32*5=21474836480 bytes (21.47484 GB).

2) The most efficient approach using built-in structures would be a slice []bool, where the index is the int representation of the IP and the value of its occurrence.
since bool uses 1 byte in Go. 2^32*1=4294967296 bytes (4.294967 GB)

3) We can take this a step further and make our code more efficient. The idea is to use only one bit to represent the occurred IP, rather than the entire byte. Unfortunately, Go does not have a built-in bit type. But bitsets come to the rescue. 2^32/1=536870912 bytes (0.5368709 GB)