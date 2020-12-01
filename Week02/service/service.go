package service

import (
	"fmt"
	"github.com/heping2018/Go-000/tree/master/model"
)

func DemandUser(id int) {
	product, err := model.ModelProduct(id)
	fmt.Println(product, err)
	fmt.Println(fmt.Sprintf("err=%+v", err))
}
