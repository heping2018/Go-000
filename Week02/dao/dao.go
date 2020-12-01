package dao

import "github.com/pkg/errors"

type Product struct {
	id   int
	name string
}

var Esql = errors.New("sql.ErrNoRows")

func QueryProduct(id int) (Product, error) {
	var p Product
	/**
	 * 查询操作
	 */
	err := errors.Wrap(Esql, "not found")
	return p, err
}
