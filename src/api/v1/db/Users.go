package db

import "api/v1/db/Schema"

func SearchUserList() []Schema.User {
	var result []Schema.User
	stms, err := DB.Prepare("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	rows, err := stms.Query()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var user_id int64
		var username string
		var password string
		err = rows.Scan(&user_id, &username, &password)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, Schema.User{user_id, username, password})
	}
	rows.Close()
	stms.Close()
	return result
}

func SearchUserById(id int64) Schema.User {
	stms, err := DB.Prepare("SELECT * FROM users WHERE user_id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var user_id int64
	var username string
	var password string

	err = row.Scan(&user_id, &username, &password)
	stms.Close()
	return Schema.User{user_id, username, password}
}

func DeleteUserById(id int64) (int64, error) {
	stms, err := DB.Prepare("delete from users WHERE user_id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(id)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	stms.Close()
	return result.RowsAffected()
}

func AddUser(user Schema.User) (int64, error) {
	stms, err := DB.Prepare("insert into users(username, password) values(?, ?)")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(user.Username, user.Password)
	stms.Close()
	return result.LastInsertId()
}

func UpdateUser(user Schema.User) (int64, error) {
	stms, err := DB.Prepare("update user set username=?,password=? where user_id=?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	result, err := stms.Exec(user.Username, user.Password, user.User_id)
	stms.Close()
	return result.RowsAffected()
}

