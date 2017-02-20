package model

import (
	"github.com/astaxie/beego/orm"

)


type Abstract interface {
	Appcreate(m *App)(id int64,err error)
        Appupdate(m *App)(id int64,err error)
	FindGuid(appguid string) (v *App, err error)
	Findall() (apps *[]*App ,err error)

}

type App struct {


	Id         int    `orm:"column(id);auto" json:"id"`
	Name       string   `orm:"column(name);unique;size(25)" json:"name"`
	Buildpack string  `orm:"column(buildpack);size(25)" json:"buildpack"`
	Instancecount int `orm:"column(instancecount)" json:"instancecount"`
	Appguid  string   `orm:"column(appguid);size(25)" json:"appguid"`
	Status  string     `orm:"column(appstatus);size(25);null" json:"status"`
	Upload_bits string `orm:"column(upload_bits);size(25);null" json:"upload_bits"`
	App_upload string   `orm:"column(app_upload);size(25);null" json:"app_upload"`
	Buildid   string    `orm:"column(buildid);size(25);null" json:"buildid"`


}

type Build struct {

	Endpoint string `json:"endpoint"`
        AccessKeyID string `json:"accesskeyid"`
        SecretAccessKey string `json:"secretaccesskey"`
        BucketName string `json:"bucketname"`
        UseSSL bool `json:"usessl"`

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

	o := orm.NewOrm()
	v = &App{}
	o.QueryTable("app").Filter("Appguid",appguid).One(v)
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (*App)Findall() (apps *[]*App ,err error){

	apps = new([]*App)
	o := orm.NewOrm()
	_,err=o.QueryTable("app").All(apps)

        return


}




