package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]
	fo, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	out := io.MultiWriter(os.Stdout, fo)
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		if _, err := out.Write(buf); err != nil {
			log.Fatal(err)
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
