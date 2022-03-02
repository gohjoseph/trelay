package main

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
