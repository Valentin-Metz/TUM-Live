package tools

import "os"

var Cfg Config

func init() {
	Cfg = Config{
		MailUser:         os.Getenv("MAIL_USER"),
		MailPassword:     os.Getenv("MAIL_PASSWORD"),
		MailServer:       os.Getenv("MAIL_SERVER"),
		DatabaseUser:     os.Getenv("MYSQL_USER"),
		DatabasePassword: os.Getenv("MYSQL_PASSWORD"),
		DatabaseName:     os.Getenv("MYSQL_DATABASE"),
		VersionTag:       os.Getenv("VERSION_TAG"),
		LrzServerIngest:  os.Getenv("LRZ_SERVER_INGEST"),
		LrzServerHls:     os.Getenv("LRZ_SERVER_HLS"),
		LrzUser:          os.Getenv("LRZ_USER"),
		LrzPassword:      os.Getenv("LRZ_PASSWORD"),
		CampusBase:       os.Getenv("CAMPUS_API_BASE"),
		CampusToken:      os.Getenv("CAMPUS_API_TOKEN"),
	}
}

type Config struct {
	MailUser         string
	MailPassword     string
	MailServer       string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	VersionTag       string
	LrzServerIngest  string
	LrzServerHls     string
	LrzUser          string
	LrzPassword      string
	CampusBase       string
	CampusToken      string
}
