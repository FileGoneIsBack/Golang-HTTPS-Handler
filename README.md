# Golang-HTTPS-Handler
This is a complex golang https/http handler

buy a domain
```
-namecheap
-cloudflare
-ovh
```

open a cloudflare account to gen your key and cert
```
link domain name servers there should be 2
```

update dns
```
-add record
--add A record pointing domain to server
```

finish the setup in cloudflare
```
update ssl/tls encryption mode to full

go to ssl/tls > origin server
-create certificate (dont close)

copy cert and key to assets cert.pem and key.pem
```

update config.go to http or https with ur keys and your good to go! 
if anyone has any recommendations on ddos protection lmk, this src can handle i think 100k rps before it crashes which is okay 

you can expand from this src ive used this before as a template. if you use htmx and make some api endpoints and add a db ur good to go its golang so pretty easy to understand and keep secure. feel free to give me credits tho ;)


