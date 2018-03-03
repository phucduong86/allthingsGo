package main

type user struct {
	username     string
	password     string
	psResetToken string
	tokenExpire  string
}

func main() {
	const expireTimeInMinute = 60

	//fake user data
	fakeuser := user{
		username:      "AuntBetty",
		password:      "secret",
		psResetToken:  "",
		timeGenerated: "",
		tokenExpired:  true,
	}

}
func resetpwHandler() {
	//handle the POST to /resetpw?email=emailaddress:
	// - get email param
	//return error if user doesn't exist
	// - generate a payload and call Generate token to has this payload, set the time + 1 day
	// payload can be time + username to make sure it is unique
}

func userResetHandler() {
	// handle GET to //userreset?username=<username>&token=...
	// get username and token param from the request
	// validateRequest
	//err OR
	//render page
}

func updatePasswordHandler() {
	// handle POST to //userreset?username=<username>&token=...
	// get username and token param from the request
	// validateRequest again
	//err OR
	//Update database with new password (Patch), expire the token
}
func GenerateToken(username string) string {
	// check last Generated time, if < 5 mins then not allowed
	// if > 5 mins, allow to:
	// - GenerateToken should return 3 values: pwResetToken string, generated time and tokenExpired=false
	// - generate link to /userreset?username=<username>&token=<pw_reset_token> and send email to client
}

func ValidateRequest(email string, token string) {
	//email exist?
	//match user token?
	//check if expired
	//validate expiry date (should be < generatedTime + 60), if > then set expired= true
	//if all checked out, extract username and do handleUpdatePassword()
}
