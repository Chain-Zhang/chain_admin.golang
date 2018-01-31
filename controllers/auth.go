package controllers

type AuthController struct{
	BaseController
}

func (self *AuthController)List(){
	self.display()
}