## iperf3
A TCP, UDP, and SCTP network bandwidth measurement tool

### Install
```shell
git clone https://github.com/esnet/iperf.git
cd iperf
./configure; make; sudo make install

sudo echo "/usr/local/lib" > /etc/ld.so.conf

```

### Use
```shell
# Server 
iperf3 -p 9000 -s 

# client
 iperf3  -c 127.0.0.1 -p 9000  -u
 
```