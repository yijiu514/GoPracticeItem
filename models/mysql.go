package models

import (
	"fmt"
)

//QueryMaxID 找到数据库当前最大id
func QueryMaxID ()int{
	var max int
	err := DB.QueryRow("select max(id) from user").Scan(&max)
	if err != nil{
		fmt.Println("query failed")
		return 0
	}
	fmt.Println("query success")
	fmt.Printf("the max id is : %d\n",max)
	return max
}

//QueryID 根据email查询id
func QueryID(email string)(int){
	var  id int
	sqlStr := "select id from user where email = ?"
	err := DB.QueryRow(sqlStr,email).Scan(&id)
	if err != nil{
		fmt.Println("query failed")
	}
	fmt.Println("query success")
	return id
}

//QuerForEditor 查询角色
func QuerForEditor(id int )(string,error){
	var role string
	sqlStr := "select role from user where id = ?"
	err := DB.QueryRow(sqlStr,id).Scan(
		&role,
	)
	if err != nil{
		fmt.Println("mysql query failed")
		return  "",err
	}
	fmt.Println("mysql query success")
	return role,nil
}

//QueryForLogin 查询登陆需要的信息
func QueryForLogin(str string)(string,string,int64){
	var u User
	sqlStr := "select password,password_salt,lock_at from user where email = ?"
	err := DB.QueryRow(sqlStr,str).Scan(
		&u.password,
		&u.passwordsalt,
		&u.lockat,
	)
	if err != nil{
		fmt.Println("query mysql failed")
	}
	return u.password,u.passwordsalt,u.lockat
}

//QuerForIdentity 查询身份相关信息
func QuerForIdentity(id int)(string,int64,string){
	var u User

	sqlStr := "select email,creat_at,role from user where id = ?"
	err := DB.QueryRow(sqlStr,id).Scan(
		&u.email,
		&u.creatat,
		&u.role,
	)
	if err != nil{
		fmt.Println("query mysql failed")
		fmt.Println(err)
	}
	fmt.Println("query mysql success")
	return u.email,u.creatat,u.role
}

// QuerySessionSalt 根据id查询sessionsalt
func QuerySessionSalt(id int)(string){
	var  salt string
	sqlStr := "select sessionsalt from user where id = ?"
	err := DB.QueryRow(sqlStr,id).Scan(&salt)
	if err != nil{
		fmt.Println("query failed")
		fmt.Println(err)
		return "query failed"
	}
	fmt.Println("query success")
	return salt
}

//InsertForRegister 插入注册信息
func InsertForRegister(email string,password string,passwordsalt string,time int64,sessionsalt string)error{
	sqlStr := "insert into user(id,email,role,creat_at,lock_at,password,passwordalt,sessionsalt) value(?,?,'editor',?,0,?,?,?)"
	_, err := DB.Exec(sqlStr, QueryMaxID()+1,email,time,password,passwordsalt,sessionsalt)
	if err != nil{

		fmt.Println(err)
		fmt.Println("insert failed")

		return  err
	}
	fmt.Println("insert success")
	return  nil
}

//UpdateForLock 更新锁定时间
func UpdateForLock(id int,time int64)error{
	sqlStr := "update user set lock_at = ? where id = ?"
	ret,err := DB.Exec(sqlStr,time,id)
	if err != nil{
		fmt.Printf("mysql update failed")
		return err
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("mysql get row wrong")
		fmt.Println(err)
		return err
	}
	fmt.Printf("mysql insert success")
	fmt.Println(n)
	return  err
}

//UpdateForPwdSet 充值密码
func UpdateForPwdSet(id string,password string,salt string) error{
	sqlStr := "update user set password = ?,password_salt = ?  where id = ?"
	ret,err := DB.Exec(sqlStr,password,salt,id)
	if err != nil{
		fmt.Printf("mysql updata flaied")
		return err
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("mysql get row wrong")
		fmt.Println(err)
		return err
	}
	fmt.Printf("mysql get row success ")
	fmt.Println(n)
	return nil
}

//UpdateForPwdChange 修改密码
func UpdateForPwdChange(id int ,password string,salt string)error{
	sqlStr := "update user set password = ?,password_salt = ?  where id = ?"
	_,err := DB.Exec(sqlStr,password,salt,id)
	if err != nil{
		fmt.Printf("msyql updata failed ")
		return err
	}
	return nil
}

//UpdateForRole 修改角色信息
func UpdateForRole(id int,role string)error {
	sqlStr := "update user set role = ? where id = ?"
	ret,err := DB.Exec(sqlStr,role,id)
	if err != nil{
		fmt.Printf("mysql update failed")
		fmt.Println(err)
		return err
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("ger row wrong")
		fmt.Println(err)
		return err
	}
	fmt.Printf("mysql insert success ")
	fmt.Println(n)
	return nil
}

//UpdateSessionsalt 更新sesscion_salt
func UpdateSessionsalt(id int,sessionsalt string)error {
	sqlStr := "update user set session _salt= ? where id = ?"
	ret,err := DB.Exec(sqlStr,sessionsalt,id)
	if err != nil{
		fmt.Printf("mysql update failed")
		fmt.Println(err)
		return err
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("ger row wrong")
		fmt.Println(err)
		return err
	}
	fmt.Printf("mysql insert success ")
	fmt.Println(n)
	return nil
}