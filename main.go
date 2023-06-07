package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	
	"sync"
	
	"github.com/syndtr/goleveldb/leveldb"
	"example.com/assignment/pkg/validation"
	
)
const NUM=1
const BNUM=2



var mutex sync.Mutex
var rwmutex  sync.RWMutex
const ( BLOCKSIZE =  5)

func main() {
	
	var txnList []validation.Txns
	var ProcessedList []validation.Txns
	validChannel := make(chan validation.Txns)
	invalidChannel:=make(chan validation.Txns)
	defer close(validChannel)
	defer close(invalidChannel)
   
	var wg sync.WaitGroup
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Print("Error while reading the data")
		return
	}

	_ = json.Unmarshal([]byte(file), &txnList)

	for i := 0; i <len(txnList); i++ {
		tempTxn := txnList[i]
		wg.Add(1)
		go validation.Push(tempTxn, validChannel,invalidChannel ,&wg)
	}

	for i := 0; i < len(txnList); i++ {
		select{
			case first:=<-validChannel:
					ProcessedList=append(ProcessedList,first)
			case second:=<-invalidChannel:
					fmt.Printf("invalid %v",second)
		}
	}
	
	//time.Sleep(3 * time.Second)
	wg.Wait()
	validation.Update(ProcessedList ,BLOCKSIZE)
	fetchBlock()
	
}
func fetchBlock(){
	var blocklist []validation.BlockData
	
	file,err:=ioutil.ReadFile("block.json")
	if err!=nil{
			fmt.Println("Error while reading the block number ")

	}
	_=json.Unmarshal([]byte(file),&blocklist)
	validation.FetchBlockNo(blocklist,2)
	fetchAllBlock()
}
func fetchAllBlock(){
	db, err := leveldb.OpenFile("./db", nil)
	defer db.Close()
	iter := db.NewIterator(nil, nil)
	
	for iter.Next() {
    	key:=iter.Key()
		value:=iter.Value()
		fmt.Println()
		fmt.Printf("Key:%vValue:%v",string(key),string(value))
		fmt.Println()
	}
	iter.Release()
	err = iter.Error()
	defer db.Close()
	fmt.Print(err)	

}
	

