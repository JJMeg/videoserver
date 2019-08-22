package dbops

import (
	"log"
	"strconv"
	"sync"
	"videoserver/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO seesions (session_id, TTL, login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL,login_name FROM seesions WHERE session_id = (?)")

	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname

	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetreiveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM seesions")
	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string

		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{
				TTL:      ttl,
				Username: login_name,
			}

			m.Store(id, ss)
		}
	}

	defer stmtOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = (?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err = stmtOut.Exec(sid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}
