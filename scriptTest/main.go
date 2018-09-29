package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type order struct {
	Id         string `json:"id"`
	UserId     string `json:"userid"`
	S          string `json:"s"`
	CreateTime string `json:"createtime"`
}

func main() {
	data, err := ioutil.ReadFile("/Users/baimengfan/dev/go/src/leetcode-alg/scriptTest/08-13.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	orders := make([]order, 0)

	err = json.Unmarshal(data, &orders)

	if err != nil {
		fmt.Println(err)
		return
	}

	res := make([]order, 0)

	for _, o := range orders {
		n, err := strconv.ParseInt(o.UserId, 10, 64)
		if err != nil {
			continue
		}
		if t := n % 10; t >= 5 {
			res = append(res, o)
		}

		if len(res) == 1000 {
			break
		}
	}

	jsonByte, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("08-13-1000.json", jsonByte, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
