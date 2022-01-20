package model

import (
	"errors"
)

// Value Object: いいね対象のタイプ
type LikeTargetType string

func CheckType(t string) bool {
	switch t {
	case "POST":
		return true
	case "COMMENT":
		return true
	default:
		return false
	}
}

func NewLikeTargetType(value string) (LikeTargetType, error) {
	if !CheckType(value) {
		return "", errors.New("targetTypeが正しくありません")
	}

	return LikeTargetType(value), nil
}
