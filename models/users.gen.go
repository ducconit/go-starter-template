// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/datatypes"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Name          string         `gorm:"column:name" json:"name"`
	Username      string         `gorm:"column:username" json:"username"`
	Email         string         `gorm:"column:email" json:"email"`
	WalletAddress string         `gorm:"column:wallet_address" json:"wallet_address"`
	Status        int8           `gorm:"column:status" json:"status"`
	Extra         datatypes.JSON `gorm:"column:extra" json:"extra"`
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
