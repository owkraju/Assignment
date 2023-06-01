package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"log"
	"os"
	"sync"
	
	"time"
	"crypto/md5"
	"github.com/syndtr/goleveldb/leveldb"
)
const NUM=1
const BNUM=2
type blockData struct {
	Blocknumber  int  `json:"blocknumber"`
	PrevBlockHash string `json:"prevhash"`
	BlockHash string `json:"blockhash"`
	Txns    []txns  `json:"txns"`
	
}
type block interface{
	push()txns
	update()txns

}

func push(tx txns, validChannel chan<- txns,invalidChannel chan<- txns, wg *sync.WaitGroup) {

	defer wg.Done()

	stat := []bool{true, false}

	
	tx.Valid = stat[rand.Int()%len(stat)]
	if  tx.Version!=1 && tx.Valid!=true{
		invalidChannel<-tx
	} else {
		validChannel <- tx
	}
	
}

var mutex sync.Mutex
var rwmutex  sync.RWMutex
const ( BLOCKSIZE =  5)
func main() {
	
	var txnList []txns
	var ProcessedList []txns
	validChannel := make(chan txns)
	invalidChannel:=make(chan txns)
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
		go push(tempTxn, validChannel,invalidChannel ,&wg)
		
	}
for i := 0; i < len(txnList); i++ {
	select{
		case first:=<-validChannel:
				ProcessedList=append(ProcessedList,first)
		case second:=<-invalidChannel:
				fmt.Printf("invalid %v",second)
		}
	}
	
	time.Sleep(3 * time.Second)
	wg.Wait()
	Committed(ProcessedList)
	
}
func Committed(pList []txns){
	
	blockfile, err := os.OpenFile("block.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	start:=0
	last:=BLOCKSIZE
	fmt.Print(len(pList))
	var previousBlock = ""
	for i:=1;i<=len(pList) ;i++{
		
		value:=(i-1)
		str:="blockchain"
		block:=&blockData{}
		if(last<len(pList)){
				startTime:=time.Now()
				
				block.Blocknumber=i
				if previousBlock != ""{
					block.PrevBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))
				}
				previousBlock = str+string(value)
				block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))

				block.Txns=pList[start:last]
				jsonData, err := json.Marshal(block)
				if err != nil {
					fmt.Println("Error:", err)
					
				}
					 if _, err := blockfile.WriteString(string(jsonData)+","); err != nil {
						log.Fatal(err)
					}
					endTime:=time.Since(startTime)
					fmt.Printf("\nTime taken to process a block number %v and time %v is\n ",i,endTime)

		}else{

					startTime:=time.Now()
					block.Blocknumber=i
					block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))
					block.PrevBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))

					block.Txns=pList[start:]
					jsonData, err := json.Marshal(block)
					if err != nil {
						fmt.Println("Error:", err)
						
					}
								 if _, err := blockfile.WriteString(string(jsonData)+","); err != nil {
									log.Fatal(err)
								}
					endTime:=time.Since(startTime)
					fmt.Printf("Time taken to process a block number %v and time %v is\n ",i,endTime)
					break
				}
				
		
		start=last
		last=last+BLOCKSIZE
		


	}
			db, err := leveldb.OpenFile("./db", nil)
			defer db.Close()

	for i:=0;i<len(pList);i++{
		var insertData txns
		insertData=pList[i]
			if insertData.Valid==true{
				mutex.Lock()
				err = db.Put([]byte(fmt.Sprintf("SIM%v\n",i)), []byte(fmt.Sprintf("%v\n",pList[i])), nil)
				mutex.Unlock()
				if err!=nil{
						fmt.Printf("error%v",err)

				
				}	
			}
	}
	 if err := blockfile.Close(); err != nil {
        log.Fatal(err)
    }
	
	fmt.Println()
	fetchBlock(db)
	}
func fetchBlock(db *leveldb.DB){
	mutex.Lock()
	data, err := db.Get([]byte(fmt.Sprintf("SIM%v\n",BNUM)), nil)
	mutex.Unlock()
	defer db.Close()
	if err != nil {
		fmt.Println("Error while getting the data:", err)
		return
	}

	fmt.Println(string(data))
	fetchAllBlock(db)
}
func fetchAllBlock(db *leveldb.DB){
	iter := db.NewIterator(nil, nil)
	
	for iter.Next() {
    	key:=iter.Key()
		value:=iter.Value()
		fmt.Println()
		fmt.Printf("Key:%vValue:%v",string(key),string(value))
		fmt.Println()
	}
	iter.Release()
	err := iter.Error()
	defer db.Close()
	fmt.Print(err)	

}
	

type txns struct{
	
	Value string `json:"value"`
	Version float32 `json:"version"`
	Valid bool `json:"valid"`
	

}