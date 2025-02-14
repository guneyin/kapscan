package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
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
