package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {

	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUser(id int) (user User, err error) {
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	user = User{}
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`

	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}

	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?`

	if err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		log.Println(err)
	}

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`

	if _, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now()); err != nil {
		log.Println(err)
	}

	session = Session{}

	cmd2 := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`

	if err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	); err != nil {
		log.Panicln(err)
	}

	return session, err
}

func (s *Session) Exists() (isValid bool, err error) {
	cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

	if err = Db.QueryRow(cmd, s.UUID).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserID,
		&s.CreatedAt,
	); err != nil {
		isValid = false
		return
	}

	if s.ID != 0 {
		isValid = true
	}

	return
}

func (s *Session) DeleteSession() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`

	_, err = Db.Exec(cmd, s.UUID)

	return err
}

func (s *Session) GetUser() (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`

	err = Db.QueryRow(cmd, s.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}
