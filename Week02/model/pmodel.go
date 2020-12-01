package model

import (
	"fmt"
	"github.com/heping2018/Go-000/tree/master/dao"
	"github.com/pkg/errors"
)

func ModelProduct(id int) (dao.Product, error) {
	product, err := dao.QueryProduct(id)
	err = errors.WithMessage(err, fmt.Sprintf("not found product id=%d ", id))
	return product, err
}
