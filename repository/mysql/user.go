package mysql

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/pkg/errmsg"
	"Q/A-GameApp/pkg/richerror"
	"database/sql"
	"time"
)

func (d *MySqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number = ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, richerror.New("mysql.IsPhoneNumberUnique").WhitWarpError(err).
			WhitMessage(errmsg.ErrorMsgCantQuery).WhitKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (d *MySqlDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number,password) value (?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, richerror.New("mysql.Register").WhitMessage("can't execute command:").WhitWarpError(err).WhitKind(richerror.KindUnexpected)

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
		return entity.User{}, false, richerror.New("mysql.GetUserByPhoneNumber").WhitMessage(errmsg.ErrorMsgCantQuery).WhitKind(richerror.KindNotFound)
	}
	return user, true, nil
}
func (d *MySqlDB) GetUserByID(userID uint) (entity.User, error) {

	row := d.db.QueryRow(`SELECT * FROM users WHERE id = ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New("mysql.GetUserByID").WhitMessage(errmsg.ErrorMsgNotFound).WhitWarpError(err).WhitKind(richerror.KindNotFound)
		}
		return entity.User{}, richerror.New("mysql.GetUserByID").WhitMessage(errmsg.ErrorMsgCantQuery).WhitWarpError(err).WhitKind(richerror.KindUnexpected)
	}
	return user, nil

}
func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt time.Time
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
