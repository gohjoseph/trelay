// Code generated by "stringer -type=ModifyTileAction"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KillTile-0]
	_ = x[PlaceTile-1]
	_ = x[KillWall-2]
	_ = x[PlaceWall-3]
	_ = x[KillTileNoItem-4]
	_ = x[PlaceWire-5]
	_ = x[KillWire-6]
	_ = x[PoundTile-7]
	_ = x[PlaceActuator-8]
	_ = x[KillActuator-9]
	_ = x[PlaceWire2-10]
	_ = x[KillWire2-11]
	_ = x[PlaceWire3-12]
	_ = x[KillWire3-13]
	_ = x[SlopeTile-14]
	_ = x[FrameTrack-15]
	_ = x[PlaceWire4-16]
	_ = x[KillWire4-17]
	_ = x[PokeLogicGate-18]
	_ = x[Actuate-19]
	_ = x[TryKillTile-20]
	_ = x[ReplaceTile-21]
	_ = x[ReplaceWall-22]
	_ = x[SlopePoundTile-23]
}

const _ModifyTileAction_name = "KillTilePlaceTileKillWallPlaceWallKillTileNoItemPlaceWireKillWirePoundTilePlaceActuatorKillActuatorPlaceWire2KillWire2PlaceWire3KillWire3SlopeTileFrameTrackPlaceWire4KillWire4PokeLogicGateActuateTryKillTileReplaceTileReplaceWallSlopePoundTile"

var _ModifyTileAction_index = [...]uint8{0, 8, 17, 25, 34, 48, 57, 65, 74, 87, 99, 109, 118, 128, 137, 146, 156, 166, 175, 188, 195, 206, 217, 228, 242}

func (i ModifyTileAction) String() string {
	if i >= ModifyTileAction(len(_ModifyTileAction_index)-1) {
		return "ModifyTileAction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ModifyTileAction_name[_ModifyTileAction_index[i]:_ModifyTileAction_index[i+1]]
}