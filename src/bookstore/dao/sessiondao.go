package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 添加Session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 删除Session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}
