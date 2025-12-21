package models

import (
	"time"
)

type Employee struct {
	Id                        int       `xorm:"not null pk autoincr comment('主键id') INT"`
	Department                string    `xorm:"comment('部门') VARCHAR(512)"`
	EmployeeID                string    `xorm:"comment('职工编码') VARCHAR(512)"`
	Name                      string    `xorm:"comment('姓名') VARCHAR(512)"`
	Position                  string    `xorm:"comment('岗位(职务)') VARCHAR(512)"`
	PositionAttribute         string    `xorm:"comment('岗位属性') VARCHAR(512)"`
	AdministrativeLevel       string    `xorm:"comment('行政级别') VARCHAR(512)"`
	Gender                    string    `xorm:"comment('性别') VARCHAR(512)"`
	Ethnicity                 string    `xorm:"comment('民族') VARCHAR(512)"`
	Age                       string    `xorm:"comment('年龄') VARCHAR(512)"`
	DateOfBirth               string    `xorm:"comment('出生日期') VARCHAR(512)"`
	StartDate                 string    `xorm:"comment('参加工作时间') VARCHAR(512)"`
	Education                 string    `xorm:"comment('学历') VARCHAR(512)"`
	ProfessionalQualification string    `xorm:"comment('专业技术职务任职资格') VARCHAR(512)"`
	QualificationDate         string    `xorm:"comment('取得时间') VARCHAR(512)"`
	IDCardNumber              string    `xorm:"comment('身份证号') VARCHAR(512)"`
	CreatedAt                 time.Time `xorm:"DATETIME"`
	UpdatedAt                 time.Time `xorm:"DATETIME"`
	DeletedAt                 time.Time `xorm:"DATETIME"`
}
