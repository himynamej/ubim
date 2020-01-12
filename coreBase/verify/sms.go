package verify


import( "net/http"
"fmt"
)
const API string ="ebCduZCcZZXQE4BH60sLzSxjz8vAwRZK"

func SendSms(mobile string ,confirm string )error  {


	_, err := http.Get("https://www.saharsms.com/api/" + API + "/json/SendVerify?receptor=" + mobile + "&template=" + "15916-confirm" + "&token=" + confirm)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	fmt.Println("code:",confirm)
	return nil

}
