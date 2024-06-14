package main

import (
	"github.com/chenlong-cxy/go-module-test/model"
	"github.com/sirupsen/logrus"
)

func main() {
	stu := model.NewStudent("Alice", 14)
	logrus.Infof("student info is %+v", stu)
}
