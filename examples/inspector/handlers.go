package main

import (
	"fmt"

	"github.com/btvoidx/trelay"
)

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

	// Reduce log spam from UpdateItemDrop packets
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

	return fmt.Sprintf("ModifyTile action: %s (%d, %d) flags: %d %d\n", action.String(), x, y, flag1, flag2)
}

// Client only
func handleConnectRequest(ctx *trelay.PacketContext, client bool) string {
	pkt := ctx.Packet()
	ver := pkt.MustReadString()

	return fmt.Sprintln("ConnectRequest", ver)
}
