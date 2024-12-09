// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

// Account
type PreGoCrmUserC struct {
	// Account ID
	UsrID uint32
	// Email
	UsrEmail string
	// Phone Number
	UsrPhone string
	// Username
	UsrUsername string
	// Password
	UsrPassword string
	// Creation Time
	UsrCreatedAt int32
	// Update Time
	UsrUpdatedAt int32
	// Creation IP
	UsrCreateIpAt string
	// Last Login Time
	UsrLastLoginAt int32
	// Last Login IP
	UsrLastLoginIpAt string
	// Login Times
	UsrLoginTimes int32
	// Status 1:enable, 0:disable, -1:delete
	UsrStatus bool
}
