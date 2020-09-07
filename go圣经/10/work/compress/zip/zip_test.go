package zip

import (
	"compress/compress"
	"fmt"
	"os"
	"testing"
)

func TestZipFile( t *testing.T) {
	src := "/tmp/a.txt"
	dst := "a_test.txt.zip"

	err := compress.Compress(dst,src)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}else {
		if err := os.Remove(dst); err != nil {
			panic(err)
		}
	}
}
