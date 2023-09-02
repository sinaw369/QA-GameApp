package mysql

import (
	"Q/A-GameApp/entity"
	"database/sql"
	"fmt"
)

func (d *MySqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can't scan query result:%w", err)
	}
	return false, nil
}
func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number,password) value (?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}
	// err is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
func (d *MySqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("can't scan query result:%w", err)
	}
	return user, true, nil
}
func (d *MySqlDB) GetUserByID(userID uint) (entity.User, error) {

	row := d.db.QueryRow(`SELECT * FROM users WHERE id = ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found:%w", err)
		}
		return entity.User{}, fmt.Errorf("can't scan query result:%w", err)
	}
	return user, nil

}
func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
