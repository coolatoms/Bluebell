package mysql

func CheckUserExist(username string) (bool, error) {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlstr, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser() {

}
