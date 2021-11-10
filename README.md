# gorestapi
In Memory Key-Value store olarak çalışan REST API projesidir.

Proje içerisinde memory-cache mekanizması mevcuttur.

Bir request Handle olduğu taktirde server log tutularak gelen istekler bir log dosyasına kayıt edilmekte.

Uygulama çalışma süresince belirlenen süre içerisinde(3600s) memory içerisindeki dosyaları "TIMESTAMP-data.gob" dosyasına kayıt eder.

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
- ### **GET** 
 In-Memory cache içerisindeki tüm dataları getirir.


- ### **GET/{Key}**  

Parametre olarak gönderilen key var ise geri getirir.

- ### **FLUSH**  

In-Memory içerisindeki tüm cache i boşaltır. TIMESTAMP-data.gob dosyasını temizler.



