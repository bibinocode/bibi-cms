/**
@Time : 2025-11-02
@Author : bibinocode
@File : enter.go
@Software: GoLand
*/

package models

import "time"

type Model struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
