package mysql

import (
	"strings"
	"studyWeb/Bluebell/models"

	"github.com/jmoiron/sqlx"
)

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content,
		post.AuthorID, post.CommunityID)
	return
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList mysql查询帖子列表函数
func GetPostList(page, size int64) (data []*models.Post, err error) {
	sqlstr := `select post_id, title, content, author_id, community_id, create_time 
	from post 
	ORDER BY create_time
	DESC 
	limit ?,?
	`
	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, sqlstr, (page-1)*size, size)
	return
}

func GetPostListByIDs(ids []string) (postlist []*models.Post, err error) {
	sqlstr := `select post_id ,title,content,author_id,community_id,create_time
	From post
	where post.post_id in(?)
    ORDER BY FIND_IN_SET(post_id,?)
`
	query, args, err := sqlx.In(sqlstr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postlist, query, args...)
	if err != nil {
		return nil, err
	}
	return
}
