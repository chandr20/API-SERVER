package controller

import ("github.com/astaxie/beego"
         "API-SERVER/model"
	"encoding/json"


	"fmt"
	"os"
	"io"
	"github.com/minio/minio-go"




	"API-SERVER/pkgs/client"




	"bytes"
	"github.com/kubernetes/client-go/pkg/util/rand"
	"API-SERVER/pkgs/kubeclient"
)



type Appcontroller struct {
	beego.Controller
}



func(app *Appcontroller)Post(){

	var Application model.App


	err:=json.Unmarshal(app.Ctx.Input.RequestBody,&Application)

	if err!=nil{
		beego.Info("Unable to unmarshal" ,err)
		app.Ctx.Output.SetStatus(500)



	}else {

		Appguid := rand.String(22)

		Application.Appguid = Appguid




	_,err = Application.Appcreate(&Application)

	if err!=nil{
		beego.Info("Unable to create app data",err)
		app.Ctx.Output.SetStatus(500)
		Application.Status = "Failed"
		app.Data["json"] = Application


	}else {
		Application.Status = "created"
		Application.Appupdate(&Application)
		app.Ctx.Output.SetStatus(200)

		app.Data["json"] = Application
	}

	}

	app.ServeJSON()



}


func (app *Appcontroller) Uploadbits() {

	var Appupload model.App

	App_data, err := Appupload.FindGuid(app.Ctx.Input.Param(":appguid"))
	if err != nil {
		beego.Info(err)
		//app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		return
	}

	file, header, err := app.GetFile("file")
	if err != nil {
		beego.Info(err)
		app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		//app.Ctx.ResponseWriter.Status(err)
		return
	}

	fileName := header.Filename
	fmt.Println(fileName)

	// save to server

	os.Mkdir("/tmp/chandra", 0777)
	pathfile := "/tmp/chandra/" + fileName
	files, err := os.Create(pathfile)
	if err != nil {
		beego.Info(err)
		app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		//app.Ctx.ResponseWriter.Status(err)
		return
	}


	// save to server
	_, err = io.Copy(files, file)

	if err != nil {
		beego.Info(err)
		app.Ctx.Output.SetStatus(400)
		//app.Ctx.ResponseWriter.Status(err)
		return
	}

	contentType := "application/txt"
	endpoint := beego.AppConfig.String("minio_endpoint")
	accessKeyID := beego.AppConfig.String("ak")
	secretAccessKey := beego.AppConfig.String("sk")
	bucketName := beego.AppConfig.String("bucketname")
	useSSL := false




	// Upload the zip file with FPutObject
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		beego.Info(err)
		app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		//app.Ctx.ResponseWriter.Status(err)
		return
	}
	_, err = minioClient.FPutObject(bucketName, fileName, pathfile, contentType)
	if err != nil {
		beego.Info(err)
		app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		//app.Ctx.ResponseWriter.Status(err)
		return
	}


	App_data.Upload_bits = fileName
	App_data.App_upload = "Upload_success"
	App_data.Appupdate(App_data)
	app.Data["json"] = App_data



	app.ServeJSON()


}




func (app *Appcontroller)Appstart(){

	App_start := new(model.App)

	App_data, err := App_start.FindGuid(app.Ctx.Input.Param(":appguid"))
	if err != nil {
		beego.Info(err)
		//app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		return
	}

	AppBuild := new(model.AppBuild)
	Build_data := new(model.Build)


	Build_data.Endpoint = beego.AppConfig.String("minio_endpoint")
	Build_data.AccessKeyID = beego.AppConfig.String("ak")
	Build_data.SecretAccessKey = beego.AppConfig.String("sk")
	Build_data.BucketName = beego.AppConfig.String("bucketname")
	Build_data.UseSSL = false

	AppBuild.App = *App_data
	AppBuild.Build = *Build_data
        //s:= fmt.Sprintf("%+v", *AppBuild)



	var cliententry client.Comm
	cliententry.Method = "POST"
	cliententry.Url = beego.AppConfig.String("buildcontroller")


	buf := new(bytes.Buffer)
        json.NewEncoder(buf).Encode(AppBuild)
	cliententry.Body = buf

	res,err := cliententry.Conn(&cliententry)
	Ap_d := new(model.App)
	json.Unmarshal(res,&Ap_d)

	Ap_d.Appupdate(Ap_d)



	if err!=nil {
		beego.Info(err)
		//app.Ctx.Output.SetStatus(400)
		app.Data["json"] = err.Error()
		app.ServeJSON()
		return
	}

	app.Data["json"] = Ap_d
	app.ServeJSON()


}

func (app *Appcontroller)Stage(){

	App_stage := new(model.AppBuild)
	json.Unmarshal(app.Ctx.Input.RequestBody,&App_stage)
	App_stage.Buildid = rand.String(22)
	Kubernetes_endpoint := beego.AppConfig.String("kube_master")
	kubenewclient := new(kubeclient.Namespace)

	err:=kubenewclient.CreateNamespace(Kubernetes_endpoint)
	if err!=nil{
		fmt.Println(err)
		app.Data["json"]= err

	}
	app.Data["json"]= App_stage


	app.ServeJSON()



}

func (app *Appcontroller)GetAll() {

	App_all := new(model.App)

	data,_ := App_all.Findall()

	app.Data["json"] = data

	app.ServeJSON()



}










