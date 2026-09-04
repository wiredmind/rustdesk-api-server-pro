package model

import "time"

type DeviceGroup struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int index"`
	Name      string    `xorm:"'name' varchar(255)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *DeviceGroup) TableName() string {
	return "device_group"
}
