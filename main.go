package main

import (
	"encoding/json"
	"github.com/chuccp/utils/log"
	core "github.com/v2fly/v2ray-core/v5"
	_ "github.com/v2fly/v2ray-core/v5/app/proxyman/inbound"
	_ "github.com/v2fly/v2ray-core/v5/app/proxyman/outbound"
	"github.com/v2fly/v2ray-core/v5/infra/conf/cfgcommon"
	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
	"os"
	"strconv"
	"time"
)

func main() {
	envport := os.Getenv("PORT")
	if envport==""{
		envport = "9096"
	}
	uuid := os.Getenv("uuid")
	num,_:=strconv.ParseUint(envport,10,32)
	var settingsJson = `{
        "udp": false,
        "clients": [
          {
            "id": "`+uuid+`",
            "alterId": 0,
            "email": "t@t.tt"
          }
        ],
        "allowTransparent": false
      }`
	settingsData := json.RawMessage(settingsJson)
	ws:=v4.TransportProtocol("ws")
	streamConfig:=v4.StreamConfig{Network:&ws,Security: "none",WSSettings:&v4.WebSocketConfig{Path: "/wowowowo"},HTTPSettings:&v4.HTTPConfig{Path: "/"}}
	inboundV4:=&v4.InboundDetourConfig{Protocol: "vmess",PortRange:&cfgcommon.PortRange{From: uint32(num), To: uint32(num)},Settings:&settingsData,StreamSetting: &streamConfig}
	vc:=&v4.Config{InboundConfigs:[]v4.InboundDetourConfig{*inboundV4},OutboundConfigs:[]v4.OutboundDetourConfig{{
		Protocol:"freedom",
	}}}
	cfg,_:=vc.Build()

	install, err := core.New(cfg)


	if err == nil {
		err = install.Start()
		log.Info("121212", err)
	} else {
		log.Info("5345345", err)
	}
	log.Info("!!!!!!", err)

	time.Sleep(time.Hour)

}
