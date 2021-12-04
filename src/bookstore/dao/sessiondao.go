package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
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

// 获取Session
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	// 执行
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	// 执行
	row := inStmt.QueryRow(sessID)
	sess := &model.Session{}
	// 扫描数据中的字段值为Session的字段赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

// 判断用户是否已经登陆
func IsLogin(r *http.Request) (bool, *model.Session) {
	// 根据Cookie的name获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 已经登陆
		// 获取Cookie的Value
		cookieValue := cookie.Value
		// 根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			return true, session
		}
	}
	// 没有登录
	return false, nil
}
