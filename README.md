# gorestapi

It is an api project based on the steps of an e-commerce site in the basket.

## Endpoints 
**POST**
### **/api/v1/basket** 
   
```javascript
{
  "guid": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "customerId": 1,
  "addressId": 1,
  "orderId": 1,
  "productList": [
    {
      "productId": 1,
      "count": 1
    }
  ]
}
```
**GET**
- ### **api/v1/basket/{basketId}** 




