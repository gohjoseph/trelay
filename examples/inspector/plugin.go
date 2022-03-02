package main

import (
	"fmt"

	"github.com/btvoidx/trelay"
)

func GetInspectorPlugin(s *trelay.Server) trelay.Plugin {
	return &plugin{s}
}

type plugin struct{ s *trelay.Server }

var filterPackets = map[trelay.PacketType]bool{
	trelay.ConnectRequest:    true,
	trelay.Disconnect:        true,
	trelay.SendTileSquare:    true,
	trelay.ModifyTile:        true,
	trelay.UpdateItemDrop:    true,
	trelay.ProjectileUpdate:  true,
	trelay.DestroyProjectile: true,
}

func (*plugin) Name() string                   { return "InspectorPlugin" }
func (*plugin) OnServerStart()                 { fmt.Println("InspectorPlugin.OnServerStart") }
func (*plugin) OnServerStop()                  { fmt.Println("InspectorPlugin.OnServerStop") }
func (*plugin) OnSessionOpen(*trelay.Session)  { fmt.Println("InspectorPlugin.OnSessionOpen") }
func (*plugin) OnSessionClose(*trelay.Session) { fmt.Println("InspectorPlugin.OnSessionClose") }

func (*plugin) OnClientPacket(ctx *trelay.PacketContext) {
	if !filterPackets[ctx.Packet().Type()] {
		return
	}

	typ := ctx.Packet().Type()
	str := "Client "
	switch typ {
	case trelay.ConnectRequest:
		str += handleConnectRequest(ctx, true)
	case trelay.SendTileSquare:
		str += handleSendTileSquare(ctx, true)
	case trelay.ModifyTile:
		str += handleModifyTile(ctx, true)
	case trelay.UpdateItemDrop:
		str += handleUpdateItemDrop(ctx, true)
	case trelay.ProjectileUpdate:
		str += handleProjectileUpdate(ctx, true)
	case trelay.DestroyProjectile:
		str += handleDestroyProjectile(ctx, true)
	default:
		str += fmt.Sprintln(typ, ctx.Packet().Length())
	}
	fmt.Print(str)
}

func (*plugin) OnServerPacket(ctx *trelay.PacketContext) {
	if !filterPackets[ctx.Packet().Type()] {
		return
	}

	typ := ctx.Packet().Type()
	str := "Server "
	switch typ {
	// case trelay.ConnectRequest:
	// 	handleConnectRequest(ctx)
	case trelay.SendTileSquare:
		str += handleSendTileSquare(ctx, false)
	case trelay.ModifyTile:
		str += handleModifyTile(ctx, false)
	case trelay.UpdateItemDrop:
		str += handleUpdateItemDrop(ctx, false)
		// Reduce log spam from UpdateItemDrop packets
		if str == "Server " {
			return
		}
	case trelay.ProjectileUpdate:
		str += handleProjectileUpdate(ctx, false)
	case trelay.DestroyProjectile:
		str += handleDestroyProjectile(ctx, false)
	default:
		str += fmt.Sprintln(typ, ctx.Packet().Length())
	}
	fmt.Print(str)
}
