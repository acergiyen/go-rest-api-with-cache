# gorestapi
It is a REST API project that works as an In Memory Key-Value store.

There is a memory-cache mechanism in the project.

In case of a request Handle, server logs are kept and incoming requests are recorded in a log file.

The application saves the files in the memory to the "TIMESTAMP-data.gob" file within the specified time (3600s) during operation.

If the project is desired to be tested in the local environment, the relevant port should be determined in the local environment or the request should be made using the "localhost/" domain by default.

Deployed on Heroku.

https://go-rest-api-with-cache.herokuapp.com/

## Endpoints 
- ### **HEADER** 
 Key : "Content/Type" , Value:"application/json"
 
- ### **SET** 
   Json/Body
   
```javascript
{
    "Key" : "1",
    "Value" :"ahmet can"         
}
```

https://go-rest-api-with-cache.herokuapp.com/api/set
- ### **GET** 
Retrieves all data in the In-Memory cache.


- ### **GET/{Key}**  

If there is a key sent as a parameter, it returns it.

https://go-rest-api-with-cache.herokuapp.com/api/get/{key}
- ### **FLUSH**  

Empties all cache in In-Memory. Clears the TIMESTAMP-data.gob file.

https://go-rest-api-with-cache.herokuapp.com/api/flush



