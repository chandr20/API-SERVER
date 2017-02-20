package client
import(

	"net/http"

	"io/ioutil"



	"fmt"


	"io"
)


type Comm struct {

	Method string

	Url string


	Body io.Reader
}


func (*Comm)Conn(s *Comm)(body []byte ,err error){




	req,err := http.NewRequest(s.Method,s.Url+"/v2/apps/stage",s.Body)
	if err!=nil{
		fmt.Println(err)
		return
	}

	res, err:= http.DefaultClient.Do(req)

	if err!=nil{
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body,err = ioutil.ReadAll(res.Body)

	return

}
