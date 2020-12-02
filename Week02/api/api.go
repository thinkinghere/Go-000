package main


import (
	"../service"
	"fmt"

)

func main()  {
	name, err := service.GetDaoUserInfo(1)
	if err != nil {
		return
	}
	fmt.Println("username: %v", name)
}
