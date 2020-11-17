package api

type DataBase interface {
	SignUp(username, password string) error
	Login(username, password string) error
	AddLink(originLink, shortenLink,userName string) error
	GetLink(shortenLink string) string
	GetUrls(username string) map[string]string
	AddSession(uuidCode, username string) error
	GetSessionInfo(uuidCode string) string
	Logout(uuidCode string) error
}
