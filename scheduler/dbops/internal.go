package dbops

import "log"

// api->videoid->db
// dispathcer -> mysql -> videoid->datachannel
// executor -> datachannel -> videoid -> delete videos

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")

	var ids []string

	if err != nil {
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord Error:%v", err)
		return ids, nil
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id=?")

	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord Error:%v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
