# Specture
## Simple QR Code + Hidden webhook utility  

###

### Run specture through
We can do some sneaky stuff with haproxy. 

Assume you have 
- A domain name: notgoogle.com
- A valid TLS/SSL key&certificate for: notgoogle.com
- An accessible server with DNS record set correctly

#### run spectrue via cmdline
> ./spectrue --listen=127.0.0.1:3000  \
-p "some_long_random_preshared_key" \
--url https://notgoogle.com \
--whitelist=/etc/haproxy/whitelist \
--dummy=https://google.com

#### Create the whitelist file
> touch /etc/haproxy/whitelist

#### Edit your /etc/haproxy/haproxy.cfg and start Haproxy

    frontend proxy-frontend
      bind *:443 transparent ssl crt /etc/haproxy/ssl.crt
        acl whitelist_local src -f /etc/haproxy/whitelist
        use_backend local_proxy if whitelist
        default_backend camouflage

    backend proxy-frontend
        server proxy-frontend 127.0.0.1:8118

    backend camouflage
        mode http
        server specture 127.0.0.1:3000

When someone visits https://notgoogle.com, they will be redirected to the actual https://google.com as a camouflage mechanism. However, if you scan the secret QR code, it will act as a webhook server, adding your external IP to the HAProxy ACL file. Now you have the access to the real backend application hidden behind HAProxy through https://notapple.com