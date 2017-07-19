package gateway

type Request struct {
	AppId      string      `json:"app_id"`
	Method     string      `json:"method"`
	Version    string      `json:"version"`
	SignType   string      `json:"sign_type"`
	Sign       string      `json:"sign"`
	Token      string      `json:"token"`
	Timestamp  string      `json:"timestamp"`
	BizContent interface{} `json:"biz_content"`
}
