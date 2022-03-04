package main

import (
	"context"
	"github.com/v2fly/v2ray-core/v5/common/net"
	"github.com/v2fly/v2ray-core/v5/features/routing"
	"github.com/v2fly/v2ray-core/v5/transport/internet"

)

type HttpInbound struct {


}

func (hi* HttpInbound) Network() []net.Network {
	return []net.Network{net.Network_TCP}
}


func (hi* HttpInbound)Process(context.Context, net.Network, internet.Connection, routing.Dispatcher) error{
	return nil
}
