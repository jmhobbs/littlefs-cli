package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseDiskutilInformation(t *testing.T) {
	output := `   Device Identifier:         disk6
   Device Node:               /dev/disk6
   Whole:                     Yes
   Part of Whole:             disk6
   Device / Media Name:       1081

   Volume Name:               Not applicable (no file system)
   Mounted:                   Not applicable (no file system)
   File System:               None

   Content (IOContent):       None
   OS Can Be Installed:       No
   Media Type:                Generic
   Protocol:                  USB
   SMART Status:              Not Supported

   Disk Size:                 31.9 GB (31914983424 Bytes) (exactly 62333952 512-Byte-Units)
   Device Block Size:         512 Bytes

   Media OS Use Only:         No
   Media Read-Only:           No
   Volume Read-Only:          Not applicable (no file system)

   Device Location:           External
   Removable Media:           Removable
   Media Removal:             Software-Activated

   Solid State:               Info not available
   Virtual:                   No
`

	blocks, blockSize, err := parseDiskutilInformation([]byte(output))
	assert.NoError(t, err)
	assert.Equal(t, int64(62333952), blocks)
	assert.Equal(t, int64(512), blockSize)
}
