### 服务端

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://dashboard.heroku.com/new?template=https://github.com/chuccp/cokeV2ray) 


"outbounds":
```json   
[
{
"tag": "proxy",
"protocol": "vmess",
"settings": {
"vnext": [
{
"address": "xxxx.herokuapp.com",
"port": 443,
"users": [
{
"id": "24b4b1e1-7a89-45f6-858c-242cf53b5bdc",
"alterId": 0,
"email": "t@t.tt",
"security": "none"
}
]
}
]
},
"streamSettings": {
"network": "ws",
"security": "tls",
"tlsSettings": {
"allowInsecure": false
},
"wsSettings": {
"path": "/wowowowo"
}
},
"mux": {
"enabled": true,
"concurrency": 8
}
},
{
"tag": "direct",
"protocol": "freedom",
"settings": {}
},
{
"tag": "block",
"protocol": "blackhole",
"settings": {
"response": {
"type": "http"
}
}
}
]
```