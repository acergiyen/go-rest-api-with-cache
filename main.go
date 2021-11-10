package main

import (
	"encoding/json"
	"fmt"
	"gorestapi/cache"
	"gorestapi/helper"
	"gorestapi/middleware"
	"gorestapi/model"
	"gorestapi/util"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var c *cache.Cache
var r *mux.Router

func init() {
	initCache()
	initLoadCacheFromFile()
	initSaveFile()
	initMuxRouter()
	//I use Goroutine. Bcs goroutine works better performance than thread
	go initInterval()
}

func main() {
	port := os.Getenv("PORT")
	
	// localport := "8080"
	fmt.Println("port:", port)
	r.HandleFunc("/", Index).Methods(http.MethodGet)
	r.HandleFunc("/api/set", Set).Methods(http.MethodPost)
	r.HandleFunc("/api/get", GetAll).Methods(http.MethodGet)
	r.HandleFunc("/api/get/{key}", Get).Methods(http.MethodGet)
	r.HandleFunc("/api/flush", Flush).Methods(http.MethodGet)
    //http logs
	httpLogServer()
    // port config from heroku
	http.ListenAndServe(":"+port, r)

}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("gorestapi"))
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	key := param["key"]
	_, resp := util.Get(c, key, r.Method)
	response := helper.JsonMarshal(resp)
	w.Write(response)
	initSaveFile()
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, resp := util.GetAll(c, r.Method)
	response := helper.JsonMarshal(resp)
	w.Write(response)
	initSaveFile()
}

//Save item by key
func Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var KeyValue model.KeyValue
	err := json.NewDecoder(r.Body).Decode(&KeyValue)
	if err == nil {
		fmt.Println("/api/set run", KeyValue)
		resp := util.Set(KeyValue, c, r.Method)
		response := helper.JsonMarshal(resp)
		w.Write(response)
		initSaveFile()
	}
}

func Flush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := util.Flush(c, r.Method)
	response := helper.JsonMarshal(resp)
	w.Write(response)
	initSaveFile()
}

func initCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

//Load the cache from saved file when you can run the project
func initLoadCacheFromFile() {
	fileName := "TIMESTAMP-data.gob"
	_, err := os.Open("tmp/" + fileName)
	if err == nil {
		//Load cache same time
		util.LoadCacheFile(c, "tmp/"+fileName)
		fmt.Println("--TIMESTAMP-data.gob run")
	} else {
		fmt.Println("TIMESTAMP-data.gob load err ", err)
	}
}
func initSaveFile() {
	fileName := "TIMESTAMP-data.gob"
	r, err := os.Create("tmp/" + fileName)
	fmt.Println("init save data: ", r)
	fmt.Println("init save err: ", err)
	if err == nil {
		//Saved added data
		util.SaveCacheFile(c, "tmp/"+fileName)
		fmt.Println("TIMESTAMP-data.gob file added data")
	} else {
		fmt.Println("initSaveFile Error", err)
	}
}
func initMuxRouter() {
	r = mux.NewRouter()
}

//runs periodically every 3600 seconds
func initInterval() {
	for range time.Tick(time.Second * 3600) {
		initSaveFile()
		util.GetAll(c, "GET")
		fmt.Println("initInterval runned", time.Now())
	}
	time.Sleep(time.Second * 5)
}

func httpLogServer() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logMiddleware := middleware.NewLogMiddleware(logger)
	r.Use(logMiddleware.Func())
}
