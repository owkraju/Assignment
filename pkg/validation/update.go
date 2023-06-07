package validation
import (
	"os"
	"fmt"
	"encoding/json"
	"time"
	"crypto/md5"
	"log"
	"sync"
	
"github.com/syndtr/goleveldb/leveldb"

)
var mutex sync.Mutex
var rwmutex  sync.RWMutex
func Update(pList []Txns,BLOCKSIZE int ){
	
	blockfile, err := os.OpenFile("block.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	var (
		blockBytes []byte
		blockList []*BlockData
	)
    _, err = blockfile.Read(blockBytes)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(blockBytes,&blockList)
	blockfile.Truncate(0)
	start:=0
	last:=BLOCKSIZE
	fmt.Print(len(pList))
	var previousBlock = ""
	for i:=1;i<=len(pList) ;i++{
		
		value:=(i-1)
		str:="blockchain"
		block:=new(BlockData)
		if(last<len(pList)){
				startTime:=time.Now()
				
				block.Blocknumber=i
				if previousBlock != ""{
					block.PrevBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))
				}
				previousBlock = str+string(value)
				block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))

				block.Txns = pList[start:last]
				blockList = append(blockList,block)
				block.BlockStatus="Committed"
				block.CreationTimeStamp=time.Now()
					endTime:=time.Since(startTime)
					fmt.Printf("\nTime taken to process a block number %v and time %v is\n ",i,endTime)

		}else{

					startTime:=time.Now()
					block.Blocknumber=i
					// block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))
					// block.PrevBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))

				if previousBlock != ""{
					block.PrevBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))
				}
				previousBlock = str+string(value)
				block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(previousBlock)))

					block.Txns=pList[start:]
					blockList = append(blockList,block)
					block.BlockStatus="Committed"
					block.CreationTimeStamp=time.Now()
					// jsonData, err := json.Marshal(blockList)
					// if err != nil {
					// 	fmt.Println("Error:", err)
						
					// }
								//  if _, err := blockfile.Write(jsonData); err != nil {
								// 	log.Fatal(err)
								// }
					endTime:=time.Since(startTime)
					fmt.Printf("Time taken to process a block number %v and time %v is\n ",i,endTime)
					break
				}
				
		
		start=last
		last=last+BLOCKSIZE
		


	}
	jsonData, err := json.Marshal(blockList)

	if _, err := blockfile.Write(jsonData); err != nil {
		log.Fatal(err)
	}

			db, err := leveldb.OpenFile("./db", nil)
			defer db.Close()

			/*
			Storing the valid t 
			*/
	for i:=0;i<len(pList);i++{
		var insertData Txns
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
	

	}