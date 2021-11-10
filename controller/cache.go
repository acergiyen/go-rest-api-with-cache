package controller

import (
	"fmt"
	"gorestapi/cache"
	"gorestapi/helper"
	"gorestapi/model"
)

func GetAll(c *cache.Cache, method string) (map[string]cache.Item, *helper.Resp) {
	items := c.Items()
	res := helper.NewResponse()
	fmt.Println("Items: ", c.Items())
	res.Success = true
	res.Action = method
	res.Data = items
	res.Errors = nil

	return items, res
}

func Get(c *cache.Cache, key string, method string) ([]byte, *helper.Resp) {
	data, exist := c.Get(key)
	res := helper.NewResponse()
	if !exist {
		fmt.Println(key, " data not exist")
		res.Success = false
		res.Action = method
		res.Data = data
		res.Errors = nil
		return nil, nil
	}

	resByte, err := data.([]byte)
	if err {
		res.Success = false
		res.Action = method
		res.Data = data
		res.Errors = nil
		return nil, nil
	}

	res.Success = true
	res.Action = method
	res.Data = data
	res.Errors = nil

	fmt.Println("get method: ", data)
	return resByte, res
}

//If exist item key, this item will update
func Set(data model.KeyValue, c *cache.Cache, method string) *helper.Resp {
	err := c.Set(data.Key, data, cache.NoExpiration)
	res := helper.NewResponse()
	if err == nil {
		res.Success = true
		res.Action = method
		res.Data = data
		res.Errors = nil
	} else {
		res.Success = false
		res.Action = method
		res.Data = nil
		res.Errors = map[string]string{"Request Body": "Provided request body malformed"}
	}

	fmt.Println("set method: ", c)

	return res
}

func Flush(c *cache.Cache, method string) *helper.Resp {
	c.Flush()
	res := helper.NewResponse()
	res.Success = true
	res.Action = method
	res.Data = ""
	res.Errors = nil
	fmt.Println("flush")

	return res
}

func SaveCacheFile(c *cache.Cache, fileName string) {
	c.SaveFile(fileName)
	fmt.Println("/controller/savecachefile run")
}

func LoadCacheFile(c *cache.Cache, fileName string) {
	c.LoadFile(fileName)
	fmt.Println("/controller/LoadCacheFile run")
}
