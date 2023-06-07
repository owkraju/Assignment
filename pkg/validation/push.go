package validation

import (
	"sync"
	
	
)
func Push(tx Txns, validChannel chan<- Txns,invalidChannel chan<- Txns, wg *sync.WaitGroup) {
	defer wg.Done()
	Status := false
	if  tx.Version!=1.0 {
		tx.Valid=Status
		invalidChannel<-tx
	} else {
		tx.Valid=true
		validChannel <- tx
	}
	
}