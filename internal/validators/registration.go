package validators

type RegistrationBody struct {
	Firstname string `binding:"required" json:"firstname"`
	Lastname  string `binding:"required" json:"lastname"`
	EmailID   string `binding:"required" json:"emailId"`
	Password  string `binding:"required" json:"password"`
}
