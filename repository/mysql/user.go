package mysql

import (
	"Q/A-GameApp/entity"
	"database/sql"
	"fmt"
)

func (d *MySqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can't scan query result:%w", err)
	}
	return false, nil
}
func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number) value (?, ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}
	// err is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil
}
