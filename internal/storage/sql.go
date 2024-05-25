package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	user = `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT,
		email TEXT UNIQUE,
		token TEXT DEFAULT NULL,
		token_duration DATETIME DEFAULT NULL
	);`
	post = `
	CREATE TABLE IF NOT EXISTS post(
		idPost INTEGER PRIMARY KEY AUTOINCREMENT,
		idAuth INTEGER,
		author	TEXT,
		title TEXT,
		content TEXT, 
		category TEXT,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		createDate DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (idAuth) REFERENCES user (id)
	);`
	category = `
	CREATE TABLE IF NOT EXISTS categories (
		name VARCHAR 
	);
	DELETE FROM categories;
	INSERT INTO categories (name) VALUES
		('Golang'),
		('Java'),
		('Python'),
		('Others');`
	comment = `
	CREATE TABLE IF NOT EXISTS comment(
		idComment INTEGER PRIMARY KEY AUTOINCREMENT,
		idPost INTEGER,
		idAuth INTEGER,
		author TEXT,
		content TEXT,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		createDate text,
		FOREIGN KEY(idPost) REFERENCES post(idPost),
		FOREIGN KEY(idAuth) REFERENCES user(id)
	);`
	reaction = `
	CREATE TABLE IF NOT EXISTS reaction(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
		postId INTEGER,
		reaction int,
		FOREIGN KEY(postId) REFERENCES post(idPost),
		FOREIGN KEY(userId) REFERENCES user(id)
	);
	`
	reactionComment = `
	CREATE TABLE IF NOT EXISTS reactionComment(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
		commentId INTEGER,
		reaction int
	)
	`
)

type ConfigDb struct {
	Driver string
	Path   string
	Name   string
}

func ConfDb() *ConfigDb {
	return &ConfigDb{
		Driver: "sqlite3",
		Name:   "forum.db",
	}
}

func CreateDb(conf *ConfigDb) (*sql.DB, error) {
	db, err := sql.Open(conf.Driver, conf.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTab(db *sql.DB) error {
	tables := []string{user, post, category, comment, reaction, reactionComment}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil

}
