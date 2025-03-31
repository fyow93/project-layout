// Package model 包含领域模型。
// 这是领域模型层，定义了领域对象。
// 上一层：领域服务层
// 下一层：无（这是最底层）

package model

type User struct {
	ID    string
	Name  string
	Email string
}
