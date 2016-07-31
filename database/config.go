package database

// Config provides all information to run a leeroy instance.
type Config struct {
	URL            string `json:"url"`
	Port           int    `json:"port"`
	Secret         string `json:"secret"`
	SSLCert        string `json:"ssl_cert"`
	SSLKey         string `json:"ssl_key"`
	ParallelBuilds int    `json:"parallel_builds"`
}
