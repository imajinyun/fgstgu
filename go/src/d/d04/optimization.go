package d04

import (
	"encoding/json"
	"strconv"
	"strings"
)

func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}
	request := Request{"demo_transaction", payload}
	v, err := json.Marshal(&request)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func processRequestWithEasyJSON(requests []string) []string {
	responses := []string{}
	for _, req := range requests {
		request := &Request{}
		request.UnmarshalJSON([]byte(req))

		var buf strings.Builder
		for _, e := range request.PayLoad {
			buf.WriteString(strconv.Itoa(e))
			buf.WriteString(",")
		}
		response := &Response{request.TransactionID, buf.String()}
		marshal, err := response.MarshalJSON()
		if err != nil {
			panic(err)
		}
		responses = append(responses, string(marshal))
	}
	return responses
}

func processRequestWithBuiltinJSON(requests []string) []string {
	responses := []string{}
	for _, req := range requests {
		request := &Request{}
		json.Unmarshal([]byte(req), request)
		ret := ""
		for _, e := range request.PayLoad {
			ret += strconv.Itoa(e) + ","
		}
		response := &Response{request.TransactionID, ret}
		marshal, err := json.Marshal(&response)
		if err != nil {
			panic(err)
		}
		responses = append(responses, string(marshal))
	}
	return responses
}
