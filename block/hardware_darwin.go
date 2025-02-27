package block

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

var diskutilInfoRegexp *regexp.Regexp = regexp.MustCompile(`^\s+(.*):\s+(.*)$`)
var diskutilSizeRegexp *regexp.Regexp = regexp.MustCompile((`\(exactly (\d+) `))

func OpenHardwareDevice(path string) (Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "diskutil", "information", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return &HardwareDevice{}, err
	}

	blocks, blockSize, err := parseDiskutilInformation(out)
	if err != nil {
		return nil, err
	}

	fd, err := OpenFileDevice(path, blockSize, blocks)
	if err != nil {
		return nil, err
	}

	return &HardwareDevice{
		FileDevice: fd,
	}, nil
}

func parseDiskutilInformation(in []byte) (int64, int64, error) {
	var blocks, blockSize int64

	scanner := bufio.NewScanner(bytes.NewReader(in))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matches := diskutilInfoRegexp.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		if matches[1] == "Device Identifier" {
			// noop
		} else if matches[1] == "Disk Size" {
			blockMatches := diskutilSizeRegexp.FindStringSubmatch(matches[2])
			if blockMatches == nil {
				continue
			}
			blocks, _ = strconv.ParseInt(blockMatches[1], 10, 64)
		} else if matches[1] == "Device Block Size" {
			fmt.Sscanf(matches[2], "%d Bytes", &blockSize)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return blocks, blockSize, nil
}
