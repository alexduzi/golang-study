package model

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Geo
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type User struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Address  Address `json:"address"`
	Company  Company `json:"company"`
}

type UserHighPost struct {
	Id    int32
	Name  string
	Posts int32
}

type Post struct {
	UserID int32
	Id     int32
	Title  string
	Body   string
}
