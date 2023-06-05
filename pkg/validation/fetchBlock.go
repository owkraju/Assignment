package validation
import (
    "fmt"
)
func FetchBlockNo(blocklist []BlockData,bNo int ){
	bLength:=len(blocklist) 

	if (bLength==1){
		fmt.Print("Only one")
	
	}
	low := 0
	high := bLength - 1
  
	for low <high{
       
       
		median := (low + high) / 2

		if(blocklist[low].Blocknumber < bNo) {
			low = median + 1
		}else{
			high = median - 1
		}
			if(blocklist[low].Blocknumber==bNo+1){
			    
				fmt.Println(blocklist[low])
			}
        // fmt.Print(blocklist[low].Blocknumber)
    }


}