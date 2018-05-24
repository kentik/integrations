package client

import (
	"fmt"
	"hash/fnv"

	"github.com/kentik/libkflow"
	"github.com/kentik/libkflow/api"
)

const (
	DEFAULT_SPEED = 10000
	DEFAULT_PORT  = "eth0"
)

type FlowClient struct {
	Sender          *libkflow.Sender
	SetSrcHostTags  map[string]map[string]bool
	SetDestHostTags map[string]map[string]bool
	interfaces      map[string]api.InterfaceUpdate
	idsByAlias      map[string]uint32
	doneInit        bool
	deviceMap       string
}

func NewFlowClient(client *libkflow.Sender, deviceMap string) *FlowClient {
	return &FlowClient{
		Sender:          client,
		deviceMap:       deviceMap,
		SetSrcHostTags:  map[string]map[string]bool{},
		SetDestHostTags: map[string]map[string]bool{},
		idsByAlias: map[string]uint32{
			"":           1, // Unknown -> eth0
			DEFAULT_PORT: 1, // Internals -> eth0
		},
		interfaces: map[string]api.InterfaceUpdate{
			DEFAULT_PORT: api.InterfaceUpdate{ // Pre-populate this with eth0 for now for external traffic
				Index:   1,
				Desc:    DEFAULT_PORT,
				Alias:   "",
				Address: "127.0.0.1",
				Speed:   DEFAULT_SPEED,
			},
		},
	}
}

func (c *FlowClient) GetInterfaceID(host string) uint32 {
	if c.deviceMap == "vmname" {
		return c.idsByAlias[DEFAULT_PORT]
	}

	if id, ok := c.idsByAlias[host]; ok {
		return id
	} else {
		return c.idsByAlias[DEFAULT_PORT] // Known vm, but not on this host, so we send out the default interface.
	}
}

func (c *FlowClient) UpdateInterfaces(isFromInterfaceUpdate bool) error {

	// Only run from not interfaces once
	if c.doneInit && !isFromInterfaceUpdate {
		return nil
	}
	c.doneInit = true

	client := c.Sender.GetClient()
	if client != nil {
		err := client.UpdateInterfacesDirectly(c.Sender.Device, c.interfaces)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FlowClient) AddInterface(intf *api.InterfaceUpdate) {

	// Interface id is defined by hash on alias.
	h := fnv.New32a()
	h.Write([]byte(intf.Alias))
	interfaceId := h.Sum32()

	intf.Index = uint64(interfaceId)
	c.idsByAlias[intf.Alias] = interfaceId
	intf.Desc = fmt.Sprintf("%d", interfaceId)
	intf.Speed = DEFAULT_SPEED

	c.interfaces[intf.Desc] = *intf
}
