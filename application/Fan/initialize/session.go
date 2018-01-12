package initialize

import "github.com/astaxie/beego/session"

func InitSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "fansessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ := session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}
