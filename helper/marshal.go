package helper

import "encoding/json"

func JsonMarshal(resp *Resp) []byte {
	response, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	return response
}