package kubeclient

import (
	"net/http"
	"fmt"
	"encoding/json"

	"bytes"
	"io/ioutil"
)

const (

	apiVersion  = "v1"
	namespace  = "chandraspace"

)

type Namespacer interface {
	CheckNamespace(namespace,kubernetes_endpoint string)(value bool,err error)
	CreateNamespace()
	Namereqbody()
}


type Namespaceres struct{

	Code int `json:"code`
}


type Metadata struct {
	Name string `json:"name"`

}


type Namespace struct {

           ApiVersion string `json:"apiVersion"`
	   Kind string `json:"kind"`
	   Metadata Metadata `json:"metadata"`

}



func (*Namespace)CheckNamespace(namespace_check *Namespace ,kubernetes_endpoint string)(value bool,err error){

	req,err := http.NewRequest("GET",kubernetes_endpoint+"/api/v1/namespaces/"+namespace,nil)
	if err!=nil{
		fmt.Println("err1",err)
		return true,err
	}

	res, err:= http.DefaultClient.Do(req)
	if err!=nil{
		fmt.Println("err2",err)
		return true,err
	}

	namespaceres := new(Namespaceres)

	body,_:= ioutil.ReadAll(res.Body)
	err =json.Unmarshal(body,&namespaceres)
	if err!=nil{
		fmt.Println(err)
		return true,err
	}
	if namespaceres.Code == 404 {


		return false,nil

	}else{

		fmt.Println("BAD")
		return true,nil
	}
}




func(Name *Namespace)CreateNamespace(kubernetes_endpoint string)(err error) {

	Namespacenew := &Namespace{
		 ApiVersion: apiVersion,
		 Kind: "Namespace",
		 Metadata:Metadata{
		 Name:namespace,
		},
	       }


	value, err := Name.CheckNamespace(Namespacenew, kubernetes_endpoint)
	if err != nil {
		fmt.Println(err)
		return err

	}
	fmt.Println("VALUE",value)
	if value == false {


		body, err := json.Marshal(Namespacenew)
		fmt.Println(body)
		fmt.Println(string(body))


		if err != nil {
			fmt.Println(err)
			return err
		}
		body_io := bytes.NewReader(body)
		fmt.Println("POST", kubernetes_endpoint + "/api/v1/namespaces/", body_io)

		req, err := http.NewRequest("POST", kubernetes_endpoint + "/api/v1/namespaces/", body_io)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return err

		}
		fmt.Println(res)


	}

	return nil

}