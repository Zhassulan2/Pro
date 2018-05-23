package model

//Settings настройки
type Settings struct {
	Host        string
	Port        int
	DBName      string
	User        string
	Password    string
	OAuth       string
	HTTPPort    int
	MaxPageSize int
}
