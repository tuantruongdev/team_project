package models

type PointUpdate struct {
	UpdateType  int    `json:"updateType"`
	DeviceID    string `json:"deviceID"`
	AdjustPoint int32  `json:"adjustPoint"`
}
