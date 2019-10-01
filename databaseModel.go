package main

// ServiceConfig struct is object parsed by gorm to be table name
type ServiceConfig struct {
	Name string // Service_Name is object parsed to be column name in table service

	Properties string // Service_Column is object to be as column name in table service
}