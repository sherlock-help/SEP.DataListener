package bizs

import(
  //bulit-in
  "fmt"

  //extends
  "SEP.DataListener/libs/ini"
  web "SEP.DataListener/libs"
  //"github.com/bakerstreet-club/otto"
)

type five struct {
  RootURL string
}

var (
    //set param
    oIni = ini.GetSecMap("web.five", &five{}).(*five)
)

//listen website with http
func ListenWebSite(){

    sDoc := web.QueryByURL(oIni.RootURL, map[string]string{
          "strSta" : "/UrpOnline/Home/Index/75_____",
          "orgId" : "75",
          "deptCode" : "",
          "sex" : "0",
          "date" : "",
          "page" : "1",
          "orderType" : "1",
          "orgType" : "1",
    }, map[string]string{
           "Accept":"*/*",
           "Accept-Encoding":"gzip, deflate",
           "Accept-Language":"zh-CN,zh;q=0.8",
           "Content-Length":"104",
           "Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
           "Cookie":"safedog-flow-item=; ASP.NET_SessionId=4j4h2d34kpm1wzz50dn3535g; sehr_xman=",
           "Host":"www.xmsmjk.com",
           "Origin":"http://www.xmsmjk.com",
           "Proxy-Connection":"keep-alive",
           "Referer":"http://www.xmsmjk.com/UrpOnline/Home/Index/75_____1",
           "User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
           "X-Requested-With":"XMLHttpRequest",
      })
      fmt.Println(web.GetPostDocSelect(sDoc, ".expert_div_index"))
    // oScripts := web.GoQueryByURLAndSelect(oIni.RootURL, "script")
    // var sAjaxScript string
    // if len(oScripts) > 7 {
    //     sAjaxScript = oScripts[7]
    // }
    // vm := otto.New()
    // vm.Run(sAjaxScript)
    // vm.Run(`
    //     var sRunIndex = function(){getIndexList()};
    // `)
    // if value, err := vm.Get("sRunIndex"); err == nil {
    //   if value, err := value.ToString(); err == nil {
    //     fmt.Printf("", value, err)
    //   }
    // }
}
