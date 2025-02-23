package util_test

import (
	"encoding/json"
	"fmt"
	reponse_util "jeanfo_mix/util/response"
	"testing"
)

func TestT(t *testing.T) {
	payload := reponse_util.ResponsePayload{
		Code: 2, Msg: "haha", Data: reponse_util.PaginatedData{
			Total: 30, Page: 10, Rows: []int{3, 4, 5},
		},
	}

	bs, err := json.Marshal(payload)
	fmt.Println(err)
	fmt.Println(string(bs))
}
