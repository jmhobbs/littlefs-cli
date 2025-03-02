package cmds

import "flag"

var (
	blocks    int64
	blockSize int64
)

func sharedFlags(fs *flag.FlagSet) {
	fs.Int64Var(&blocks, "blocks", 0, "littlefs block count, 0 is auto-detect")
	fs.Int64Var(&blockSize, "block-size", 0, "littlefs block size, 0 is auto-detect for devices, or 512 for images")
}

var commonFlagSet *flag.FlagSet = flag.NewFlagSet("littlefs", flag.ExitOnError)

func init() {
	sharedFlags(commonFlagSet)
}
