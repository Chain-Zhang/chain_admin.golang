package controllers

type IndexController struct{
	BaseController
}

func (self *IndexController) Index(){
    self.display()
}