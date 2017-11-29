package main

import (
	"fmt"
	// "github.com/jackshapow/shapow/api/database"
	// "github.com/labstack/echo"
	// middleware "github.com/labstack/echo/middleware"
	// // "github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	// "github.com/jackshapow/shapow/api/model"
	"github.com/dgraph-io/badger"
	// // "time"
	//"github.com/dgraph-io/badger"
	//"github.com/jackshapow/shapow/api/controller"
	// "github.com/jackshapow/shapow/api/model/pb"
	//"github.com/golang/protobuf/proto"
	"reflect"
	// "strconv"
	"github.com/davecgh/go-spew/spew"
)

func test() {
	opt := badger.DefaultOptions
	opt.Dir = "database/badger"
	opt.ValueDir = "database/badger"
	db, err := badger.Open(opt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.TypeOf(db))

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)

		//		it := txn.NewIterator(&DefaultIteratorOptions)
		prefix := []byte("u:")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.Value()
			if err != nil {
				return err
			}
			fmt.Printf("%s", k)
			spew.Dump(v)
		}
		return nil
	})

	// u := model.User{Id: "e98u219e8u3e", Name: "Bobby Jenkins", Email: "bjenkins@gmail.comz", Password: "passwordhere"}

	// data, err := proto.Marshal(&u)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Encoded data: ")
	// fmt.Println(data)

	// key := []byte("u:" + u.Id)

	// fmt.Println("Store in badger...")
	// // for r := 1; r <= 1; r++ {

	// for i := 1; i <= 10; i++ {
	// 	nkey := []byte("u:" + u.Id + string(i))

	// 	// nkey := key + string(i)
	// 	// strconv.FormatInt(n, 2)
	// 	fmt.Println("Stores " + string(nkey))
	// 	KV.Set(nkey, data, 0x00)
	// }

	// // }

	// KV.RunValueLogGC(0.5)

	// var item badger.KVItem

	// if err := KV.Get(key, &item); err != nil {
	// 	fmt.Printf("Error while getting key: %q", key)
	// 	return
	// }

	// var val []byte
	// item.Value(func(v []byte) error {
	// 	val = make([]byte, len(v))
	// 	copy(val, v)
	// 	return nil
	// })

	// fmt.Println("Fetched item:")
	// fmt.Println(val)

	// fmt.Println("Now unmarshall...")

	// var newu model.User

	// if err := proto.Unmarshal(val, &newu); err != nil {
	// 	fmt.Println("Failed to parse user:", err)
	// }

	// fmt.Println(newu)
	//fmt.Println(reflect.TypeOf(newu))

	//  if err := model.User.Unmarshal(item, &u); err != nil {
	//  	fmt.Println("Failed to parse user:", err)
	//  }

	// fmt.Println(u)

	// opt := badger.DefaultOptions
	// //dir, _ := ioutil.TempDir("", "badger")
	// dir := "database/badger"

	// opt.Dir = dir
	// opt.ValueDir = dir
	// kv, _ := badger.NewKV(&opt)

	// key := []byte("hello")

	// kv.Set(key, []byte("world"), 0x00)
	// fmt.Printf("SET %s world\n", key)

	// var item badger.KVItem
	// if err := kv.Get(key, &item); err != nil {
	//   fmt.Printf("Error while getting key: %q", key)
	//   return
	// }
	// var val []byte
	// err := item.Value(func(v []byte) error {
	//   val = make([]byte, len(v))
	//   copy(val, v)
	//   return nil
	// })
	// if err != nil {
	//   fmt.Printf("Error while getting value for key: %q", key)
	//   return
	// }

	// fmt.Printf("GET %s %s\n", key, val)

	// if err := kv.CompareAndSet(key, []byte("venus"), 100); err != nil {
	//  fmt.Println("CAS counter mismatch")
	// } else {
	//  if err = kv.Get(key, &item); err != nil {
	//    fmt.Printf("Error while getting key: %q", key)
	//  }

	//  err := item.Value(func(v []byte) error {
	//    val = make([]byte, len(v))
	//    copy(val, v)
	//    return nil
	//  })

	//  if err != nil {
	//    fmt.Printf("Error while getting value for key: %q", key)
	//    return
	//  }

	//  fmt.Printf("Set to %s\n", val)
	// }
	// if err := kv.CompareAndSet(key, []byte("mars"), item.Counter()); err == nil {
	//  fmt.Println("Set to mars")
	// } else {
	//  fmt.Printf("Unsuccessful write. Got error: %v\n", err)
	// }
}
