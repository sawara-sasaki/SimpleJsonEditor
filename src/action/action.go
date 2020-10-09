package action

import (
	"fmt"
  "net/http"
  "encoding/json"
)

type ActionRequest struct {
		Action string        `json:"action"`
		Parameters []float64 `json:"parameters"`
}

type ActionResponse struct {
		Status int  `json:"status"`
		Data	[]int `json:"data"`
}

func Handle(request []byte)(ActionResponse, error) {
	var req ActionRequest
	var res ActionResponse
	var err error
	json.Unmarshal(request, &req)
	switch req.Action {
	case "linear":
		if len(req.Parameters) == 3 {
			res.Data = GetLinearFloatData(req.Parameters[0], req.Parameters[1], req.Parameters[2],)
		} else if len(req.Parameters) == 2 {
			res.Data = GetLinearIntData(int(req.Parameters[0]), int(req.Parameters[1]))
		} else {
			err = fmt.Errorf("err %s", "Bad Parameters")
		}
	default:
		err = fmt.Errorf("err %s", "Bad Action")
	}
	if err == nil {
		res.Status = http.StatusOK
	}
	return res, err
}

func GetLinearIntData(start int, end int) []int {
	var res []int
	for i := start; i <= end; i++ {
		res = res.append(i)
	}
	return res
}

func GetLinearFloatData(start float64, end float64, length int) []float64 {
	var res []float64
	delta := (end - start) / float64(length)
	for i := start; i <= end; i += delta {
		res = res.append(i)
	}
	return res
}
