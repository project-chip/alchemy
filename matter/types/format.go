package types

type NumberFormat uint8

const (
	NumberFormatUndefined NumberFormat = iota
	NumberFormatInt
	NumberFormatHex
	NumberFormatAuto
)
