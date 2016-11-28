package main

import (
	"crypto/md5"
	"crypto/sha1"
	"io"
	"log"
)

func main() {
	{
		t := md5.New()
		io.WriteString(t, "11")
		log.Printf("md5 = %x", t.Sum(nil))
	}
	{
		t := sha1.New()
		io.WriteString(t, "11")
		log.Printf("sha1 =%x", t.Sum(nil))
	}

}
