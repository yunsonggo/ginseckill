# ginseckill
This is a E-commerce seckill system based on gin.

go run main.go

go run validate.go

go run getOne.go

postman:
post: 0.0.0.0:8090/manager/sign/up

json: 
{
    "name":"name",
    "password":"password",
}
set the cookie of user

get: 0.0.0.0:8081/check?productID=1

go run consumer.go
