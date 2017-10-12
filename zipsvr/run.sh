#!/usr/bin/env bash
docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/sunho/go/src/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
sunhok2/zipserver
