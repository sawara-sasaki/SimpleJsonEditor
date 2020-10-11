package action

import (
	"fmt"
	"math"
	"strconv"
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
			res.Data = GetLinearFloatData(req.Parameters[0], req.Parameters[1], req.Parameters[2])
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

func GetLinearFloatData(start float64, end float64, delta float64) []interface{} {
	var res []interface{}
	digitStr := "%." + fmt.Sprintf("%.0f", -1 * math.Log10(delta)) + "f"
	length := int((end - start) / delta)
	for i := 0; i < length; i++ {
		f, _ := strconv.ParseFloat(fmt.Sprintf(digitStr, start + float64(i) * delta), 64)
		res = append(res, f)
	}
	return res
}
