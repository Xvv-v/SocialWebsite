package po

//User 对象
type User struct {
	Userid       string `json:"userid,omitempty"`
	Username     string `json:"username,omitempty"`
	Telephone    string `json:"telephone,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Faceicon     string `json:"faceicon,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Region       string `json:"region,omitempty"`
	Education    string `json:"education,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

//将