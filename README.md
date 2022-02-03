## TECHNOLOGIES

- GOLANG ` 1.14 `
- GORM
- GORILLA-MUX
- POSTGRESQL 


## API's PATHS
## For Product Service 
##### For Get All Products ```GET localhost:8080/api/products ``` <br/>
##### For Add Product ```POST localhost:8080/api/products ```  <br/>
###### Request Body: 
```json
{
    "product_name":"",
    "description":"",
    "price":0.0,
    "category_id":1
}
```
##### For Update Product ```PATCH localhost:8080/api/products/update ```  <br/>
###### Request Body: 
```json
{
    "id":1,
    "product_name":"",
    "description":"",
    "price":0.0,
    "category_id":1
}
```

##### For Delete Product ```DELETE localhost:8080/api/products/{id} ```  <br/>
## For Category Service </hr>
##### For Get All Categories ```GET localhost:8080/api/categories ``` <br/>
##### For Add Category ```POST localhost:8080/api/categories ```  <br/>
###### Request Body: 
```json
{
    "category_name":""
}
```

##### For Update Category ```PATCH localhost:8080/api/categories/update ```  <br/>
###### Request Body: 
```json
{
    "id":1,
    "category_name":""
}
```


##### For Delete Category ```DELETE localhost:8080/api/categories/{id} ```  <br/>
