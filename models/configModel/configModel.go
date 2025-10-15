package configModel

type DBConfigModel struct {
	User string
	Pass string
	Name string
	Host string
	Port string
}

type EMailConfigModel struct {
	User string
	Pass string
	SMTP string
	Port string
}

type ConfigModel struct {
	DBConf    *DBConfigModel
	EMailConf *EMailConfigModel
	URLSite   string
}
