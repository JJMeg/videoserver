package dbops

import "testing"

//init(dblogin,truncate tables)-> run test ->clear data(truncate tables)

func cleatTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	cleatTables()
	m.Run()
	cleatTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("testAddUserCredential", testAddUserCredential)
	t.Run("testGetUserCredential", testGetUserCredential)
	t.Run("testDeleteUser", testDeleteUser)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("AVE", "123")
	if err != nil {
		t.Errorf("Error of AddUserCredential: %v", err)
	}

}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("AVE")
	if err != nil || pwd != "123" {
		t.Errorf("Error of GetUserCredential: %v", err)
	}

}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("AVE", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("AVE")
	if err != nil || pwd != "" {
		t.Errorf("Error of RegetUser: %v", err)
	}
}
