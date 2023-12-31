package object // import "github.com/1Volk/vksdk/object"

// SecureLevel struct.
type SecureLevel struct {
	Level int `json:"level"` // Level
	UID   int64 `json:"uid"`   // User ID
}

// SecureSmsNotification struct.
type SecureSmsNotification struct {
	AppID   int    `json:"app_id"`  // Application ID
	Date    int    `json:"date"`    // Date when message has been sent in Unixtime
	ID      int    `json:"id"`      // Notification ID
	Message string `json:"message"` // Message text
	UserID  int64    `json:"user_id"` // User ID
}

// SecureTokenChecked struct.
type SecureTokenChecked struct {
	Date    int `json:"date"`    // Date when access_token has been generated in Unixtime
	Expire  int `json:"expire"`  // Date when access_token will expire in Unixtime
	Success int `json:"success"` // Returns if successfully processed
	UserID  int64 `json:"user_id"` // User ID
}

// SecureTransaction struct.
type SecureTransaction struct {
	Date    int `json:"date"`     // Transaction date in Unixtime
	ID      int `json:"id"`       // Transaction ID
	UIDFrom int64 `json:"uid_from"` // From ID
	UIDTo   int64 `json:"uid_to"`   // To ID
	Votes   int `json:"votes"`    // Votes number
}
