package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnection struct.
type DbConnection struct {
	Dsn string
	Db  *sql.DB
}

// UserTable struct.
type UserTable struct {
	UID       int
	Email     string
	Nickname  string
	Status    uint8
	CreatedAt string
	UpdatedAt string
}

func main() {
	var dropTable string = `DROP TABLE IF EXISTS user;`
	var createTable string = `
	CREATE TABLE user (
	id int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	email varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
	nickname varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
	status tinyint unsigned DEFAULT '1' COMMENT '状态(1: 正常, 0: 禁用)',
	created_at datetime NOT NULL COMMENT '创建时间',
	updated_at datetime NOT NULL COMMENT '更新时间',
	PRIMARY KEY (id),
	UNIQUE KEY email (email) COMMENT '邮箱'
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';`
	var err error
	dsn := "root:12300123@tcp(127.0.0.1:3306)/test?charset=utf8"
	db := DbConnection{Dsn: dsn}
	db.Db, err = sql.Open("mysql", db.Dsn)
	defer db.Db.Close()
	if err != nil {
		panic(err)
	}
	_, err = db.Db.Exec(dropTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✔️ Drop user table successfully!")
	_, err = db.Db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("✔️ Create user table successfully!")

	insertData(&db)
	fmt.Println("✔️ General mode insert data successfully!")

	preInsertData(&db)
	fmt.Println("✔️ Prepare mode insert data successfully!")

	user := db.QueryRowData("SELECT * FROM user ORDER BY id DESC LIMIT 1")
	fmt.Println("✔️ Query one result:", user)

	users := db.QueryAllData("SELECT * FROM user")
	fmt.Println("✔️ Query all result:")
	for _, v := range users {
		fmt.Println("  ", v)
	}

	users = db.PreQueryAllData("SELECT * FROM user WHERE id<? ORDER BY created_at DESC", 5)
	fmt.Println("✔️ Query all result:")
	for _, v := range users {
		fmt.Println("  ", v)
	}
}

// ExecSQL func.
func (db *DbConnection) ExecSQL(sql string) (count, id int64, err error) {
	result, err := db.Db.Exec(sql)
	if err != nil {
		panic(err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
	count, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, id, nil
}

// PreInsertData func.
func (db *DbConnection) PreInsertData(sql string, args ...interface{}) (count, id int64, err error) {
	stmt, err := db.Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	if id, err = res.LastInsertId(); err != nil {
		panic(err)
	}
	if count, err = res.RowsAffected(); err != nil {
		panic(err)
	}
	return count, id, nil
}

// PreQueryAllData func.
func (db *DbConnection) PreQueryAllData(sql string, args ...interface{}) (users map[int]UserTable) {
	stmt, err := db.Db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	users = make(map[int]UserTable)
	entity := new(UserTable)
	for rows.Next() {
		err := rows.Scan(
			&entity.UID,
			&entity.Email,
			&entity.Nickname,
			&entity.Status,
			&entity.CreatedAt,
			&entity.UpdatedAt)
		if err != nil {
			panic(err)
		}
		users[entity.UID] = *entity
	}
	return users
}

// QueryAllData func.
func (db *DbConnection) QueryAllData(sql string) (user map[int]UserTable) {
	rows, err := db.Db.Query(sql)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	user = make(map[int]UserTable)
	entity := new(UserTable)
	for rows.Next() {
		err := rows.Scan(
			&entity.UID,
			&entity.Email,
			&entity.Nickname,
			&entity.Status,
			&entity.CreatedAt,
			&entity.UpdatedAt)
		if err != nil {
			panic(err)
		}
		user[entity.UID] = *entity
	}
	return user
}

// QueryRowData func.
func (db *DbConnection) QueryRowData(sql string) (data UserTable) {
	user := new(UserTable)
	err := db.Db.QueryRow(sql).Scan(
		&user.UID,
		&user.Email,
		&user.Nickname,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return *user
}

func insertData(db *DbConnection) {
	at := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("  Datetime: ", at)
	sql := "INSERT INTO user SET email='test@qq.com',nickname='test'," +
		"created_at='" + at + "',updated_at='" + at + "'"
	fmt.Println("  SQL: ", sql)
	count, id, err := db.ExecSQL(sql)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("  Affected rows: ", count)
		fmt.Println("  Last inserted id: ", id)
	}
}

func preInsertData(db *DbConnection) {
	at := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO user(email,nickname,created_at,updated_at) VALUES(?,?,?,?)"
	fmt.Println("  SQL: ", sql)
	count, id, err := db.PreInsertData(sql, "demo@qq.com", "demo", at, at)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("  Affected rows: ", count)
		fmt.Println("  Last inserted id: ", id)
	}
}
