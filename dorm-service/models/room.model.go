package models

type Room struct {
	SquareFoot   float32    `json:"squareFoot"`
	Toalet       ToaletType `json:"toalet"`
	NumberOfBeds int        `json:"numberOfBeds"`
}
