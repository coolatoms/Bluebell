package mysql

import (
	"studyWeb/Bluebell/models"
	"studyWeb/Bluebell/setting"
	"testing"
)

func init() {
	conf := setting.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&conf)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	p := &models.Post{
		ID:          1,
		AuthorID:    123,
		CommunityID: 1,
		Status:      0,
		Title:       "test",
		Content:     "just a test",
	}

	err := CreatePost(p)
	if err != nil {
		t.Fatalf("createPost insert record into mysql failed,err;%v\n", err)

	}
	t.Logf("createPost insert record into mysql success")
}
