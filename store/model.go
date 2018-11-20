package store

// User is a collector user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtToken is a token
type JwtToken struct {
	Token string `json:"token"`
}

// Exception is an exception
type Exception struct {
	Message string `json:"message"`
}

// MacAddress represents a mac Address access point
type MacAddress struct {
	ID           int    `bson:"_id"`
	placeId      string `json:"placeId"`
	mac          string `json:"mac"`
	ssid         string `json:"ssid"`
	lat          string `json:"lat"`
	lng          string `json:"lng"`
	creationDate string `json:"creationDate"`
	createdBy    string `json:"createdBy"`
	validated    bool   `json:"validated"`
}

// MacAddresses is an array of MacAddress objects
type MacAddresses []MacAddress

// Scan represents a scan
type Scan struct {
	ID           int          `bson:"_id"`
	creationDate string       `json:"creationDate"`
	createdBy    string       `json:"createdBy"`
	checked      bool         `json:"checked"`
	macList      MacAddresses `json:"macList"`
}

// Scans is an array of Scan objects
type Scans []Scan
