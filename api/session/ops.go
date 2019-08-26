package session

import (
	"log"
	"sync"
	"time"

	"videoserver/api/dbops"
	"videoserver/api/defs"
	"videoserver/api/utils"
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
	id, _ := utils.NewUUID()
	ct := newInMilli()
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{
		Username: id,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := newInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}

func newInMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
