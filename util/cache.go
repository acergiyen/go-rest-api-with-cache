package	util

import (
	"gorestapi/cache"
	"gorestapi/controller"
	"gorestapi/helper"
	"gorestapi/model"
)
// get single item by key
func Get(c *cache.Cache, key string, method string) ([]byte, *helper.Resp) {
	return controller.Get(c, key, method)

}
// get all items
func GetAll(c *cache.Cache, method string) (map[string]cache.Item, *helper.Resp) {
	return controller.GetAll(c, method)
}
// set item 
func Set(keyVal model.KeyValue, c *cache.Cache, method string) *helper.Resp {
	return controller.Set(keyVal, c, method)
}

// clear cache
func Flush(c *cache.Cache, method string) *helper.Resp {
	return controller.Flush(c, method)
}
//save cache file
func SaveCacheFile(c *cache.Cache, fileName string) {
	controller.SaveCacheFile(c, fileName)
}

func LoadCacheFile(c *cache.Cache, fileName string) {
	controller.LoadCacheFile(c, fileName)
}