package main

import (
	"fmt"
	// "github.com/google/uuid"
	"io/ioutil"
	"os"
	"strconv"
	// // "log"
	"math/rand"
	"encoding/json"
	"path/filepath"
	"example.com/assignment/pkg/validation"
)


func transcation(count int) validation.Txns {
	
	stat := []string{"1.0","2.0"}

	fmt.Print(stat[rand.Int()%len(stat)])
	 
	var token validation.Txns
	token.Value=count
	token.Version=stat[rand.Int()%len(stat)]
	return token

}
func main() {
	
	var txnList []validation.Txns

	args := os.Args[1]
	trans, _ := strconv.Atoi(args)
	for i := 1; i <=trans; i++ {
		
		data := transcation(i)
		txnList = append(txnList, data)
	

	}
	path := filepath.Join("./", "test.json")
	
	data_mar, _ := json.Marshal(txnList)
	fmt.Print("length is ", len(txnList))
	err := ioutil.WriteFile(path, data_mar, 0644)
	if err != nil {
		fmt.Print("Not happening")

	}
	
}
