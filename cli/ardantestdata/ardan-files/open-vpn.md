This instructions are incomplete but I wanted to document IP fowarding.

## Create EC2

Create an OpenVPN server in AWS.

https://docs.openvpn.net/getting-started/amazon-web-services-ec2-tiered-appliance-quick-start-guide/

## Connect to Server

The original `pem` is required to SSH into OpenVPN server.

```
chmod 400 LupaVPCSQLServer.pem
ssh -i "LupaVPCSQLServer.pem" root@35.168.165.45
```

## Route Public IP to Private IP

```
echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -I PREROUTING -d 34.197.192.162 -j DNAT --to-destination 10.0.196.44
iptables -t nat -I POSTROUTING -j MASQUERADE
```

After you are confident you haven't locked yourself out of the server save.

```
iptables-save
```
