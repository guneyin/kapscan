package util

import (
	"strconv"
	"strings"
	"time"
)

type VersionInfo struct {
	Version    string
	CommitHash string
	BuildTime  string
}

var (
	Version    string
	CommitHash string
	BuildTime  string

	lastRun time.Time
)

func GetVersion() *VersionInfo {
	return &VersionInfo{
		Version:    Version,
		CommitHash: CommitHash,
		BuildTime:  BuildTime,
	}
}

func SetLastRun(t time.Time) {
	lastRun = t
}

func GetLastRun() time.Time {
	return lastRun
}

type Money struct {
	amount string
}

func NewMoney(amount string) *Money {
	return &Money{
		amount: strings.TrimSpace(amount),
	}
}

func (m *Money) Float64() float64 {
	dec, _ := strconv.ParseFloat(m.amount, 64)
	return dec
}
