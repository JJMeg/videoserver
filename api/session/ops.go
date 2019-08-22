package session

import (
	"log"
	"sync"
	"videoserver/api/dbops"
	"videoserver/api/defs"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.RetreiveAllSessions()
	if err != nil {
		log.Printf("%s", err)
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})

}

func GenerateNewSessionId(un string) string {
	return ""
}

func IsSessionExpired(sid string) (string, bool) {
	return "", false
}
