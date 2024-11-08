package model

const TableNameGoCrmUserV2 = "go_crm_user_v2"

// GoCrmUserV2 Account
type GoCrmUserV2 struct {
	UsrID          int32  `gorm:"column:usr_id;primaryKey;autoIncrement:true;comment:Account ID" json:"usr_id"`                       // Account ID
	UsrEmail       string `gorm:"column:usr_email;not null;comment:Email" json:"usr_email"`                                           // Email
	UsrPhone       string `gorm:"column:usr_phone;not null;comment:Phone Number" json:"usr_phone"`                                    // Phone Number
	UsrUsername    string `gorm:"column:usr_username;not null;comment:Username" json:"usr_username"`                                  // Username
	UsrPassword    string `gorm:"column:usr_password;not null;comment:Password" json:"usr_password"`                                  // Password
	UsrCreatedAt   int32  `gorm:"column:usr_created_at;not null;comment:Creation Time" json:"usr_created_at"`                         // Creation Time
	UsrUpdatedAt   int32  `gorm:"column:usr_updated_at;not null;comment:Update Time" json:"usr_updated_at"`                           // Update Time
	UsrCreateIP    string `gorm:"column:usr_create_ip;not null;comment:Creation IP" json:"usr_create_ip"`                             // Creation IP
	UsrLastLoginAt int32  `gorm:"column:usr_last_login_at;not null;comment:Last Login Time" json:"usr_last_login_at"`                 // Last Login Time
	UsrLastLoginIP string `gorm:"column:usr_last_login_ip;not null;comment:Last Login IP" json:"usr_last_login_ip"`                   // Last Login IP
	UsrLoginTimes  int32  `gorm:"column:usr_login_times;not null;comment:Login Times" json:"usr_login_times"`                         // Login Times
	UsrStatus      bool   `gorm:"column:usr_status;not null;comment:Status (1: enabled, 0: disabled, -1: deleted)" json:"usr_status"` // Status (1: enabled, 0: disabled, -1: deleted)
}

// TableName GoCrmUserV2's table name
func (*GoCrmUserV2) TableName() string {
	return TableNameGoCrmUserV2
}
