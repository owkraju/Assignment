package validation
import 
(
	"time"

)
type Txns struct{
	
	Value int `json:"value"`
	Version string `json:"version"`
	Valid bool `json:"valid"`
	

}


type BlockData struct {
	Blocknumber  int  `json:"blocknumber"`
	PrevBlockHash string `json:"prevhash"`
	BlockHash string `json:"blockhash"`
	Txns    []Txns  `json:"txns"`
	CreationTimeStamp time.Time `json:"creationtime"`
	BlockStatus string `json:"bstatus"`
	
}
type block interface{
	Push()Txns
	Update()Txns

}