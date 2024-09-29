# Counting the number of unique IP addresses in a very large file

You have a simple text file with IPv4 addresses. One line is one address, line by line:
```
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
```
The file is unlimited in size and can occupy tens and hundreds of gigabytes.

You should calculate the number of unique addresses in this file using as little memory and time as possible. There is a "naive" algorithm for solving this problem (read line by line, put lines into HashSet). It's better if your implementation is more complicated and faster than this naive algorithm.

## Solutions

In theory, there can be 2^32 unique IP addresses.

1) Using map[uint32]bool to store the occurred IP addresses would be inefficient.
since uint32 uses 4 bytes and bool uses 1 byte in go. 2^32*5=21474836480 bytes (21.47484 GB).

2) The most efficient approach using built-in structures would be a slice []bool, where the index is the int representation of the IP and the value of its occurrence.
since bool uses 1 byte in Go. 2^32*1=4294967296 bytes (4.294967 GB)

3) We can take this a step further and make our code more efficient. The idea is to use only one bit to represent the occurred IP, rather than the entire byte. Unfortunately, Go does not have a built-in bit type. But bitsets come to the rescue. 2^32/1=536870912 bytes (0.5368709 GB)