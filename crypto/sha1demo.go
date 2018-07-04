package main

import (
	"crypto/sha1"
	"fmt"
)

/**
CONTEXT:
SHA1 hashes are frequently used to compute short identities for binary or text blobs.
e.g., the git revision control system uses SHA1s extensively to identify versioned files
and directories.
 */
func main()  {
	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
