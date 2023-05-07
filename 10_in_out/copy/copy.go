package copy

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var source string
var dest string
var offset int
var limit int

func init() {
	flag.StringVar(&source, "from", "", "file to read from")
	flag.StringVar(&dest, "to", "", "file to write to")
	flag.IntVar(&offset, "offset", 0, "offset in input file")
	flag.IntVar(&limit, "limit", 0, "limit in input file")
}

// Copy copies files using chunks.
func Copy() {
	flag.Parse()

	if limit < offset {
		println("Offset must be less than limit.")
		return
	}
	sourceObj, err := os.OpenFile(source, os.O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Panicf("file does not exist: %v", err)
		}
		// другие ошибки, например нет прав
	}
	defer sourceObj.Close()
	sourceObj.Seek(int64(offset), io.SeekStart)

	destObj, err := os.Create(dest)
	if err != nil {
		if os.IsExist(err) {
			log.Panicf("file already exists: %v", err)
		}
	}
	defer destObj.Close()

	done := 0
	n := 1024
	for done < limit-offset {
		if (done + n) > limit-offset {
			n = limit - offset - done
		}
		written, err := io.CopyN(destObj, sourceObj, int64(n))
		done += n
		fmt.Printf("Done: %d\n", done)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("End of file, written: %d\n", written)
				break
			}
			log.Panicf("can't copy: %v", err)
		}
		fmt.Printf("written : %d\n", written)
	}
}
