package entity

type RefFormat struct {
	FormatId         string `gorm:"type:VARCHAR(255)" json:"format_id"`
	Format           string `gorm:"type:VARCHAR(255)" json:"format"`
	Counter          int    `gorm:"type:INT" json:"counter"`
	CounterPadLength int    `gorm:"type:INT" json:"counter_pad_length"`
	CounterPadChar   string `gorm:"type:CHAR(1)" json:"counter_pad_char"`
	ExtraField       string `gorm:"type:VARCHAR(255)" json:"extra_field"`
}

func (RefFormat) TableName() string {
	return "ref_format"
}
