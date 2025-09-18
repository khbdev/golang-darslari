package storage

import "mikroservice/model"

var Users []model.User
var Orders []model.Order
var Products []model.Product  

func AddUser(u model.User) {
    Users = append(Users, u)
}

func GetUsers() []model.User {
    return Users
}


func AddOrder(o model.Order) {
    Orders = append(Orders, o)
}

func GetOrders() []model.Order {
    return Orders
}


func AddProduct(p model.Product) {
    Products = append(Products, p)
}

func GetProducts() []model.Product {
    return Products
}
