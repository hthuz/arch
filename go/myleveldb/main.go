package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	db1, err := leveldb.OpenFile("./data", nil)
	if err != nil {
		panic(err)
	}

	db2, err := leveldb.OpenFile("./data", nil)
	if err != nil {
		panic(err)
	}

	db1.Put([]byte("key1"), []byte("1"), nil)
	db2.Put([]byte("key2"), []byte("2"), nil)
	db1.Put([]byte("key3"), []byte("3"), nil)
	db2.Put([]byte("key4"), []byte("4"), nil)
	db2.Put([]byte("key1"), []byte("5"), nil)

	v, _ := db1.Get([]byte("key1"), nil)
	fmt.Println(string(v))
	v, _ = db2.Get([]byte("key2"), nil)
	fmt.Println(string(v))
	v, _ = db1.Get([]byte("key3"), nil)
	fmt.Println(string(v))
	v, _ = db2.Get([]byte("key4"), nil)
	fmt.Println(string(v))
	v, _ = db2.Get([]byte("key1"), nil)
	fmt.Println(string(v))
	db1.Close()
	db2.Close()

}
