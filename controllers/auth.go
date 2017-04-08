package controllers

import (
	"fmt"
	"my_go_web/models"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

//sign up
func (c *AuthController) GetSignup() {

}
func (c *AuthController) PostSignup() {

}

//sign in
func (c *AuthController) GetLogin() {

}
func (c *AuthController) PostLogin() {
	email := c.GetString("email")
	user, err := models.GetUsersByEmail(email)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名不存在"}
	} else {
		//比较密码
		//string to []byte
		password := []byte(c.GetString("password"))
		//http://stackoverflow.com/questions/23259586/bcrypt-password-hashing-in-golang-compatible-with-node-js

		// Hashing the password with the default cost of 10  DefaultCost int = 10
		//laravel bcrypt /Library/WebServer/Documents/estate/vendor/laravel/framework/src/Illuminate/Hashing/BcryptHasher.php
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(hashedPassword))

		// Comparing the password with the hash
		db_hashed_password := []byte(user.Password)
		err = bcrypt.CompareHashAndPassword(db_hashed_password, password)
		if err == nil {// nil means it is a match
			c.Data["json"] = map[string]interface{}{"code":1, "message": "登陆成功"}
		}
	}
	c.ServeJSON()

}

//find password 填写email
func (c *AuthController) GetResetPassword() {

}
func (c *AuthController) PostResetPassword() {
	//获取email地址 发送邮件

}