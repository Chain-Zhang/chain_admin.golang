package controllers

import(

)

type ErrorController struct{
	BaseController
}

func (self *ErrorController) Error404(){
    self.display()
}

func (self *ErrorController) Error403(){
    self.display()
}