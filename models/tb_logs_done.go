package models

import "time"

type TbLogsSyncDoneMtani struct {
	ID      int       `gorm:"column:id;" foreignKey:"id"`
	TTable  string    `gorm:"column:t_table"`
	TId     int       `gorm:"column:t_id"`
	TglInp  time.Time `gorm:"column:tgl_inp"`
	TStatus int       `gorm:"column:t_status"`
}
