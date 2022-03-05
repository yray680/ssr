**因为v2ray 中vmess是较为广泛的协议，支持的设备也多，个人感觉没必要弄太多协议**
### 服务端

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://dashboard.heroku.com/new?template=https://github.com/chuccp/cokeV2ray) 

### 客户端
"outbounds":
```json   
  "outbounds": [
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
      }
    }
  ]
```