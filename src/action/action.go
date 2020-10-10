package action

import (
	"fmt"
	"encoding/json"
)

type ActionRequest struct {
	Action     string    `json:"action"`
	Parameters []float64 `json:"parameters"`
}

type ActionResponse struct {
	Status int           `json:"status"`
	Data   []interface{} `json:"data"`
}

func Handle(request []byte)(ActionResponse, error) {
	var req ActionRequest
	var res ActionResponse
	var err error
	json.Unmarshal(request, &req)
	switch req.Action {
	case "linear":
		if len(req.Parameters) == 3 {
			res.Data = GetLinearFloatData(req.Parameters[0], req.Parameters[1], int(req.Parameters[2]))
		} else if len(req.Parameters) == 2 {
			res.Data = GetLinearIntData(int(req.Parameters[0]), int(req.Parameters[1]))
		} else {
			err = fmt.Errorf("err %s", "Bad Parameters")
		}
	default:
		err = fmt.Errorf("err %s", "Bad Action")
	}
	return res, err
}

func GetLinearIntData(start int, end int) []interface{} {
	var res []interface{}
	for i := start; i <= end; i++ {
		res = append(res, i)
	}
	return res
}

func GetLinearFloatData(start float64, end float64, length int) []interface{} {
	var res []interface{}
	delta := (end - start) / float64(length)
	for i := start; i <= end; i += delta {
		res = append(res, i)
	}
	return res
}
