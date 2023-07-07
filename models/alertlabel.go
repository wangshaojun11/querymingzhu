package models

type AlertLabel struct {
	Value    string `gorm:"column:Value"`
	GroupKey string `gorm:"column:groupKey"`
}

// 定义AlertLabel模型的表名
func (AlertLabel) TableName() string {
	return "AlertLabel"
}
