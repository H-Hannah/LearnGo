// 使用 Gob 传输数据，与json/xml不同
// Gob 并不是一种不同于 Go 的语言，而是在编码和解码过程中用到了 Go 的反射。
// Gob 使用通用的 io.Writer 接口，通过 NewEncoder() 函数创建 Encoder 对象并调用 Encode()；
// 相反的过程使用通用的 io.Reader 接口，通过 NewDecoder() 函数创建 Decoder 对象并调用 Decode()

// gob1.go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3, 4, 5, "python"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
}

// Go 中的加密
// hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；
// crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。
// 以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。
