# Specture
## Sneaky webhook utility works with HAProxy

     -c, --config=    Config file
     -p, --key=       Preshared key
     -e, --expire=    QR Credential expire time (default: 0, never expire)
     --listen=    Listen address (default: 0.0.0.0:3000)
     --log-level= Log level (default: info)
     --url=       Your websites' base url
     --dummy=     Dummy url (default: https://www.baidu.com)
     --whitelist= Location of HAProxy ACL file (default: /etc/haproxy/whitelist)


### Run specture through

Assume you have 
- A domain name, for example: notgoogle.com
- A valid TLS/SSL key&certificate for: notgoogle.com
- An accessible server with DNS record set correctly

#### Create the whitelist file
> touch /etc/haproxy/whitelist

#### Edit your /etc/haproxy/haproxy.cfg and start Haproxy

    frontend proxy-frontend
      bind *:443 transparent ssl crt /etc/haproxy/ssl.crt
        acl whitelist_local src -f /etc/haproxy/whitelist
        use_backend local_proxy if whitelist
        default_backend camouflage
        option  forwardfor

    backend proxy-frontend
        server proxy-frontend 127.0.0.1:8118

    backend camouflage
        mode http
        server specture 127.0.0.1:3000
        option  forwardfor
#### run spectrue via cmdline
> ./spectrue --listen=127.0.0.1:3000  \
-p "some_long_random_preshared_key" \
--url https://notgoogle.com \
--whitelist=/etc/haproxy/whitelist \
--dummy=https://google.com


When someone visits https://notgoogle.com, they will be redirected to the dummy url https://google.com as a camouflage mechanism. However, if you scan the secret QR code, it will act as a webhook server, adding your external IP to the HAProxy ACL file. Now you have the access to the real backend application hidden behind HAProxy through https://notapple.com

### Keynote
This whole concept relies on a secured connection, which is provided by TLS. You have to make sure you're using certificates from a trusted provider.