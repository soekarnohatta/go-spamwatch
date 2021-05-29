// Package spamwatch provides a HTTP wrapper for the SpamWatch API.
package spamwatch

// Permission is a basic int type
type Permission int

const (
	// Root is the highest role in the API
	Root Permission = iota

	// Admin can do more things than the User
	// role
	Admin

	// User is the basic and the minimum role
	// to access the API
	User
)

type (
	// SWType consists all type of SpamWatch API
	SWType struct {
		Ban
		Stats
		Token
		Version
	}

	// Ban contains all information regarding the
	// details of a banned user if any
	Ban struct {
		// Id is the banned Telegram user id
		Id      int `json:"id"`

		// Reason is the reason behind the ban of him/her
		Reason  string `json:"reason"`

		// Date is the time of the user has been banned
		Date    int `json:"date"`

		// Admin is the admin who bans the user
		Admin   int `json:"admin"`

		// Message contains text message(s) of the banned user
		Message string `json:"message"`
	}

	// Stats contains the number of banned user(s)
	Stats struct {
		// TotalBanCount is the number of banned users
		TotalBanCount int `json:"total_ban_count"`
	}

	// Token contains a detailed token information
	Token struct {
		// Id holds the id of the token
		Id         int `json:"id"`

		Permission `json:"permission"`

		// Token holds the user token
		Token      string `json:"token"`

		// UserId holds the user id of the token
		UserId     int `json:"userid"`

		// Retired holds retired status
		Retired    bool `json:"retired"`
	}

	// Version contains the detailed versioning of
	// the API
	Version struct {
		// Major is the major version of the API
		Major      string `json:"major"`

		// Minor is the minor version of the API
		Minor      string `json:"minor"`

		// Patch is the patch version of the API
		Patch      string `json:"patch"`

		// ApiVersion the full version of the API
		ApiVersion string`json:"version"`
	}
)
