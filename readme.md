# Go Lang CRUD Restful With MySQL + ORM + Gin Web Framework

Features : 
1. CRUD Restful
2. Gin Web Framework
3. MySQL ORM with [gorm.io](gorm.io)
4. Model & Migration

How to : 
1. Install Go Lang
2. Run `git mod tidy`
3. Run `git run .`
4. Base path : localhost:8080

Endpoint : 

| Name | Endpoint | Method | Params
| --- | --- | --- | --- |
| List data | /products | GET | - |
| Detail Data | /products/detail | GET | id |
| Create | /products/create | POST | sku,name,description,stock,price |
| Update | /products/update | POST | sku,name,description,stock,price |
| Delete | /products/delete | POST | id |