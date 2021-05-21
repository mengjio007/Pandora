package Pinocchio

import "github.com/skip2/go-qrcode"

type RecoveryLevel = qrcode.RecoveryLevel

const (
	// Level L: 7% error recovery.
	Low RecoveryLevel = iota

	// Level M: 15% error recovery. Good default choice.
	Medium

	// Level Q: 25% error recovery.
	High

	// Level H: 30% error recovery.
	Highest
)

func QrCodeWriteFile(content string, level RecoveryLevel, size int, filename string) {
	qrcode.WriteFile(content, level, size, filename)
}
