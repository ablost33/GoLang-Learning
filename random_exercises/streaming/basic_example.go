package streaming

import (
	"bufio"
	"io"
	"log"
	"os"
)

func StreamingPractice() {
	src := bufio.NewReader(os.Stdin)
	dst := bufio.NewWriter(os.Stdout)

	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	if err := dst.Flush(); err != nil {
		log.Fatal(err)
	}
}
