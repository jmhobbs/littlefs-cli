> [!WARNING]  
> This is a work in progress, and currently targeted at Mac OS only.

# littlefs CLI

This is a small command line tool for working with [littlefs](https://github.com/littlefs-project/littlefs) file systems.  It can work directly on block devices or on images.

# Usage

Paths for this tool are defined as `volume:path`, where `volume` can be a block device (i.e. `/dev/disk6`) or an image file (i.e. `~/image.lfs`).  The path is the path to a file on the littlefs file system.

Path is optional for some commands, and defaults to the root of the file system.

```
USAGE
  littlefs <command>

SUBCOMMANDS
  cat   Output contents of a file on a littlefs filesystem.
  cp    Copy files to, from, or on a littlefs filesystem.
  fmt   Format a block device, or create a new image file.
  ls    List files on a littlefs filesystem.
  mv    Move files to, from, or in a littlefs filesystem.
  rm    Remove files from a littlefs filesystem.
  tree  Walk a littlefs filesystem and print a tree view.

FLAGS
  -block-count 0  littlefs block count, 0 is auto-detect
  -block-size 0   littlefs block size, 0 is auto-detect for devices, or 512 for images
```

## Block Size  / Block Count

Block size and count are auto-detected where possible, but can be overriden with flags.
