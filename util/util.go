package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func deepCopy(src, dest any) (any, error) {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(src)
	if err != nil {
		return nil, err
	}
	err = gob.NewDecoder(&buf).Decode(dest)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func Convert[T any](from any, to T) (T, error) {
	res, err := deepCopy(from, to)
	if err != nil {
		return to, err
	}

	if rt, ok := res.(T); ok {
		return rt, nil
	}

	return to, fmt.Errorf("cannot convert from %T to %T", from, to)
}
