package mo

// VolunteeUserLogin table comment
type VolunteeUserLogin struct {
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
}

// TableName sets the insert table name for this struct type
func (v *VolunteeUserLogin) TableName() string {
	return "voluntee_user_login"
}
