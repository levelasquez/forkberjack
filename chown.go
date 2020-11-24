// +build !linux

package forkberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
