package bizs

import(
  //bulit-in
  "fmt"

  //extends
  "SEP.DataListener/libs/ini"
  web "SEP.DataListener/libs"
)

type Five struct {
  RootURL string

}

var (
    //set param
    oApi = ini.GetSecMap("web.five", &Five{}).(*Five)
)

//listen website with http
func ListenWebSite(){
    fmt.Println(web.QueryByURL(oApi.RootURL))
}
