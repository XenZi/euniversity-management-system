package models

import "fmt"

type ApplicationType int

const (
	Budget ApplicationType = iota
	SelfFinancing
	Disability
	SensitiveGroups
)

func (t ApplicationType) String() string {
	switch t {
	case Budget:
		return "Budget"
	case SelfFinancing:
		return "SelfFinancing"
	case Disability:
		return "Disabilty"
	case SensitiveGroups:
		return "SensitiveGroups"
	default:
		return fmt.Sprintf("Unknown TipPrijave value: %d", int(t))
	}
}

type ToaletType int

const (
	RoomShared ToaletType = iota
	FloorShared
	RoomBased
)

func (t ToaletType) String() string {
	switch t {
	case RoomShared:
		return "RoomShared"
	case FloorShared:
		return "FloorShared"
	case RoomBased:
		return "RoomBased"
	default:
		return "Unkown"
	}
}

type ApplicationStatus int

const (
	Review ApplicationStatus = iota
	Accepted
	Denied
	Pending
)
