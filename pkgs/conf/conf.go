package conf

import "time"

type PortRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type NeoBoxConf struct {
	NeoWorkDir          string        `json:"neo_work_dir"`         // dir to store files
	NeoLogFileDir       string        `json:"neo_log_dir"`          // dir to store log files
	XLogFileName        string        `json:"x_log_file_name"`      // log file name of sing-box/xray
	RawUriURL           string        `json:"download_url"`         // where to download raw proxies
	RawUriFileName      string        `json:"download_file_name"`   // file name of raw proxies
	ParsedFileName      string        `json:"parse_file_name"`      // file name of parsed proxies
	PingedFileName      string        `json:"pinged_file_name"`     // file name of ping succeeded proxies
	MaxPingers          int           `json:"max_pinger_count"`     // number of pingers
	MaxAvgRTT           int64         `json:"max_pinger_avgrtt"`    // threshold of ping avg_rtt, in milliseconds
	VerifiedFileName    string        `json:"verified_file_name"`   // file name of verification succeeded proxies
	VerifierPortRange   *PortRange    `json:"verifier_port_range"`  // number of goroutines to verify the proxies
	VerificationUri     string        `json:"verification_uri"`     // google url for verification
	VerificationTimeout time.Duration `json:"verification_timeout"` // in seconds
	VerificationCron    string        `json:"verification_cron"`    // crontab for verifier
	NeoBoxClientInPort  int           `json:"neo_box_client_in_port"`
}

func GetDefaultConf() (n *NeoBoxConf) {
	n = &NeoBoxConf{}
	n.NeoWorkDir = `C:\Users\moqsien\data\projects\go\src\neobox`
	n.NeoLogFileDir = n.NeoWorkDir
	n.XLogFileName = "neobox_xlog.log"
	n.RawUriURL = "https://gitlab.com/moqsien/xtray_resources/-/raw/main/conf.txt"
	n.RawUriFileName = "neobox_raw_proxies.json"
	n.ParsedFileName = "neobox_parsed_proxies.json"
	n.PingedFileName = "neobox_pinged_proxies.json"
	n.MaxPingers = 100
	n.MaxAvgRTT = 600
	n.VerifiedFileName = "neobox_verified_proxies.json"
	n.VerifierPortRange = &PortRange{
		Min: 4000,
		Max: 4050,
	}
	n.VerificationUri = "https://www.google.com"
	n.VerificationTimeout = 3
	n.VerificationCron = "@every 2h"
	n.NeoBoxClientInPort = 2019
	return
}
