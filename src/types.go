package main

type EmployeeResponse struct {
	Employees []string `json:"employees" xml:"employees"`
	Count     int      `json:"count" xml:"count"`
}

type ProductResponse struct {
	Products []string `json:"products" xml:"products"`
	Count    int      `json:"count" xml:"count"`
}
