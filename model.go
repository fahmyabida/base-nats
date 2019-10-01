package main

// Configuration type is used to define common configuration to get database or to get server configuration
type Configuration struct {
	IP       string      `json:"ip"`
	Port     string      `json:"port"`
	User     string      `json:"user"`
	Password string      `json:"password"`
	DBName   string      `json:"name"`
	Engine   interface{} //varible to create instance of object , this variable could be anything. Example : Engine.(*sql.Conn) => onject sql connetion
}

//Properties is used to map config file
type Properties struct {
	IP       string `json:"ip_database"`
	Port     string `json:"port_database"`
	User     string `json:"user_database"`
	Password string `json:"password_database"`
	DBName   string `json:"name_database"`
	Config   string `json:"config_name"`
	LogPath  string `json:"log_path"`
}

//Service is used to generate service object
type Service struct {
	URL    string
	Engine interface{}
}
