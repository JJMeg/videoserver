package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//init(dblogin,truncate tables)-> run test ->clear data(truncate tables)
var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate sessions")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("testAddUserCredential", testAddUserCredential)
	t.Run("testGetUserCredential", testGetUserCredential)
	t.Run("testDeleteUser", testDeleteUser)
	t.Run("testRegetUser", testRegetUser)
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

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUserCredential)
	t.Run("testAddNewVideoInfo", testAddNewVideoInfo)
	t.Run("testGetVideoInfo", testGetVideoInfo)
	t.Run("testDeleteVideoInfo", testDeleteVideoInfo)
	t.Run("testRegetVideoInfo", testRegetVideoInfo)
}

func testAddNewVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddNewVideoInfo:%v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I  like it!"
	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of Add Comments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1e9, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Errorf of List Comments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v\n", i, ele)

}
