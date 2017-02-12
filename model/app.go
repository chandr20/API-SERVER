package model

import (
	"github.com/astaxie/beego/orm"
)

type App struct {


	Id         int    `orm:"column(id);auto"`
	Name       string   `orm:"column(name);unique;size(25);null"`
	Buildpack string  `orm:"column(buildpack);size(25);null"`
	Instancecount int `orm:"column(instancecount)"`
	Appguid  string   `orm:"column(appguid);size(25);null"`
	Status  string     `orm:"column(appstatus);size(25);null"`
	Upload_bits string `orm:"column(Upload_bits);size(25);null`
	App_upload string   `orm:"column(App_upload);size(25);null`


}

type Build struct {

	Endpoint string
        AccessKeyID string
        SecretAccessKey string
        BucketName string
        UseSSL bool

}


type AppBuild struct {
	App
	Build
}




func init(){
	orm.RegisterModel(new(App))

}

func (*App)Appcreate(m *App)(id int64,err error){
	o:= orm.NewOrm()
	id, err = o.Insert(m)
	return
}


func (*App)Appupdate(m *App)(id int64,err error){
	o:= orm.NewOrm()
	id, err=o.Update(m)
	return
}


func (*App)FindGuid(appguid string) (v *App, err error) {
	//beego.Info("IMPORTANT",appguid)
	o := orm.NewOrm()
	v = &App{}
	o.QueryTable("app").Filter("Appguid",appguid).One(v)
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}



