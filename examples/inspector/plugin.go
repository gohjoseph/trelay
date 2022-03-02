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

func handleUpdateItemDrop(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()

	itemID := pkt.MustReadInt16()
	posX := pkt.MustReadFloat32()
	posY := pkt.MustReadFloat32()
	vecX := pkt.MustReadFloat32()
	vecY := pkt.MustReadFloat32()
	stackSize := pkt.MustReadInt16()
	prefix := pkt.MustReadByte()
	noDelay := pkt.MustReadByte()
	netItemID := pkt.MustReadInt16()

	if stackSize == 0 {
		return ""
	}

	return fmt.Sprintf("UpdateItemDrop id %d stack %d pos (%.2f, %.2f) vec (%.2f, %.2f) prefix %d noDelay %d netID %d\n",
		itemID, stackSize, posX, posY, vecX, vecY, prefix, noDelay, netItemID)
}

func handleProjectileUpdate(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()

	projID := pkt.MustReadInt16()
	posX := pkt.MustReadFloat32()
	posY := pkt.MustReadFloat32()
	vecX := pkt.MustReadFloat32()
	vecY := pkt.MustReadFloat32()
	owner := pkt.MustReadByte()
	typ := pkt.MustReadInt16()
	flags := pkt.MustReadByte()

	str := fmt.Sprintf("ProjUpdate id %d owner %d type %d pos (%.2f, %.2f) vec (%.2f, %.2f)", projID, owner, typ, posX, posY, vecX, vecY)

	if flags&0x01 != 0 {
		ai1 := pkt.MustReadFloat32()
		str += fmt.Sprintf(" ai1 %.2f", ai1)
	}
	if flags&0x02 != 0 {
		ai2 := pkt.MustReadFloat32()
		str += fmt.Sprintf(" ai2 %.2f", ai2)
	}
	if flags&0x08 != 0 {
		bidtrt := pkt.MustReadUint16()
		str += fmt.Sprintf(" BIDTRT %d", bidtrt)
	}
	if flags&0x10 != 0 {
		dmg := pkt.MustReadUint16()
		str += fmt.Sprintf(" dmg %d", dmg)
	}
	if flags&0x20 != 0 {
		knbk := pkt.MustReadFloat32()
		str += fmt.Sprintf(" knbk %.2f", knbk)
	}
	if flags&0x40 != 0 {
		origDmg := pkt.MustReadUint16()
		str += fmt.Sprintf(" origDmg %d", origDmg)
	}
	if flags&0x80 != 0 {
		uuid := pkt.MustReadUint16()
		str += fmt.Sprintf(" uuid %d", uuid)
	}
	return str + "\n"
}
func handleDestroyProjectile(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()

	projID := pkt.MustReadUint16()
	owner := pkt.MustReadByte()

	return fmt.Sprintf("ProjKill id %d owner %d\n", projID, owner)
}

func handleSendTileSquare(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()

	x := pkt.MustReadInt16()
	y := pkt.MustReadInt16()
	xSize := pkt.MustReadByte()
	ySize := pkt.MustReadByte()
	typ := pkt.MustReadByte()

	return fmt.Sprintf("STS (%d, %d) size (%d, %d) typ %d len %d\n", x, y, xSize, ySize, typ, pkt.Length()-3-7)
}

func handleModifyTile(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()

	action := ModifyTileAction(pkt.MustReadByte())
	x := pkt.MustReadInt16()
	y := pkt.MustReadInt16()
	flag1 := pkt.MustReadInt16()
	flag2 := pkt.MustReadByte()

	if client {
		if action == KillTile {

		}
	}

	return fmt.Sprintf("ModifyTile action: %s (%d, %d) flags: %d %d\n", action.String(), x, y, flag1, flag2)
}

// Client only
func handleConnectRequest(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()
	ver := pkt.MustReadString()

	writer := trelay.PacketWriter{}
	writer.SetType(pkt.Type())
	// if ver == "Terraria248" {
	// 	ver = "Terraria247"
	// }
	writer.WriteString(ver)

	ctx.Session().Server.Write(writer.Packet())
	ctx.SetHandled()
	return fmt.Sprintln("ConnectRequest", ver)
}

// func printPacket(ctx *trelay.PacketContext) {
// 	pkt := ctx.Packet()
// 	fmt.Println("packet", pkt.String())
// 	fmt.Println("packet", hex.EncodeToString(pkt.MustReadBytes(pkt.Length()-3)))
// }

type ModifyTileAction byte

const (
	KillTile ModifyTileAction = iota
	PlaceTile
	KillWall
	PlaceWall
	KillTileNoItem
	PlaceWire
	KillWire
	PoundTile
	PlaceActuator
	KillActuator
	PlaceWire2
	KillWire2
	PlaceWire3
	KillWire3
	SlopeTile
	FrameTrack
	PlaceWire4
	KillWire4
	PokeLogicGate
	Actuate
	TryKillTile
	ReplaceTile
	ReplaceWall
	SlopePoundTile
)
