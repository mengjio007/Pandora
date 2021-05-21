package Pinocchio

import uuid "github.com/satori/go.uuid"

type UUID = uuid.UUID

func NewUuidV1() string {
	return uuid.NewV1().String()
}

func NewUuidV2(domain byte) string {
	return uuid.NewV2(domain).String()
}

func NewUuidV3(ns UUID, name string) string {
	return uuid.NewV3(ns, name).String()
}

func NewUuidV4() string {
	return uuid.NewV4().String()
}

func NewUuidV5(ns UUID, name string) string {
	return uuid.NewV5(ns, name).String()
}
