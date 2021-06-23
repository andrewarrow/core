package lib

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteOpenDB() *sql.DB {
	db, err := sql.Open("sqlite3", "clout.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func SqliteInsertPostEntry(sdb *sql.DB, postEntry *PostEntry) {
	tx, _ := sdb.Begin()

	body := string(post.Body)
	hash := base58.Encode(post.PostHash.Bytes())
	author := base58.Encode(post.PosterPublicKey)

	s := `insert into posts (author, hash, body, created_at) values (?, ?, ?, ?)`
	thing, e := tx.Prepare(s)
	if e != nil {
		fmt.Println(e)
		return
	}
	_, e = thing.Exec(author, hash, body, time.Now())
	if e != nil {
		fmt.Println(e)
		return
	}

	e = tx.Commit()
	if e != nil {
		fmt.Println(e)
	}
}

func SqliteCreateSchema(sdb *sql.DB) {
	sqlStmt := `
create table posts (author text, hash text, body text, created_at datetime);

CREATE UNIQUE INDEX posts_hash_idx
  ON posts (hash);

CREATE INDEX posts_username_idx
  ON posts (author);

create table users (bio text, username text, pub58 text, created_at datetime);

CREATE UNIQUE INDEX users_idx
  ON users (pub58);

CREATE INDEX users_username_idx
  ON users (username);

create table user_follower (followee text, follower text);

CREATE INDEX uf_followee_idx
  ON user_follower (followee);

CREATE INDEX uf_follower_idx
  ON user_follower (follower);
`
	_, err := sdb.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
