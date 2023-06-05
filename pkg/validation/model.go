package validation

type Txns struct{
	
	Value string `json:"value"`
	Version float32 `json:"version"`
	Valid bool `json:"valid"`
	

}


type BlockData struct {
	Blocknumber  int  `json:"blocknumber"`
	PrevBlockHash string `json:"prevhash"`
	BlockHash string `json:"blockhash"`
	Txns    []Txns  `json:"txns"`
	CreationTimeStamp string `json:"creationtime"`
	BlockStatus string
	
}
type block interface{
	Push()Txns
	Update()Txns

}