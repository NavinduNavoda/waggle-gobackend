package services

type LoggedUser struct {
	Username  string
	SessionId string
}

type Auth struct {
	Users map[string]string
}

func GetNewSessionId() string {
	return "gdtug16723161##22^&%&%$#%231"
}

func CheckSessionId(sessionId string) bool {
	return sessionId == "gdtug16723161##22^&%&%$#%231"
}