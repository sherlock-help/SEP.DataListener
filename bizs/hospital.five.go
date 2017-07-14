package bizs

import(
  //bulit-in
  "fmt"
  "strings"
  "strconv"
  "time"

  //extends
  "SEP.DataListener/libs/ini"

  . "bakerstreet-club/logs"

  GoQuery "github.com/bakerstreet-club/goquery"
  web "SEP.DataListener/libs"
  "SEP.DataListener/libs"
  //"github.com/bakerstreet-club/otto"

  fiveHospistal "SEP.DataListener/domain/xmsmjk"
)

type five struct {
  QueryGap time.Duration
  RootURL string
}

var (
    //set param
    oIni = ini.GetSecMap("web.five", &five{}).(*five)

    oDoctors = []*fiveHospistal.Doctor{}
)



var oArrayDepts []string
var oArrayDoctors []string
var oMapNumbers map[string]string
func init(){

    //oArrayDepts = make([]string)
    //oArrayDoctors = make([]string)
    //统计号源
    oMapNumbers = make(map[string]string)

    //set default
    //fmt.Println(oIni.QueryGap)
    //fmt.Println("========================================")
    if 0 == oIni.QueryGap {
        oIni.QueryGap = 1000000000
    }
}

func loopPageQuery(sPage string){
    oPostResponse := web.GetPostResponse(oIni.RootURL, map[string]string{
          "strSta" : "/UrpOnline/Home/Index/75_____",
          "orgId" : "75",
          "deptCode" : "",
          "sex" : "0",
          "date" : "",
          "page" : sPage,
          "orderType" : "1",
          "orgType" : "1",
    }, map[string]string{
           "Accept":"*/*",
           //"Accept-Encoding":"gzip, deflate",
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

    oAllDom := web.GetPostDocSelection(oPostResponse, "*")

    // fmt.Println(sPage)
    // fmt.Println("++++++++++++++++++++++++++")
     //fmt.Println(oAllDom.Text())
    // fmt.Println("++++++++++++++++++++++++++")
    // fmt.Println(oAllDom.Find(".list_paging"))
    // fmt.Println("++++++++++++++++++++++++++")

    oAllDom.Find(".expert_div_index").Each(func(_ int, v *GoQuery.Selection){
      oThisWeek := v.Find(".div_index_bottom .div_bottom_isweek .isweek_ind")
      oNextWeek := v.Find(".div_index_bottom .div_bottom_nsweek .isweek_ind")


      //put in doctor
      oThisWeekInfo := []map[string]string{}
      oNextWeekInfo := []map[string]string{}

      //这周
      var sThisFullText string
      var sThisNoFullText string
      //fmt.Println(oThisWeek.Find("a"))
      var iThisWeekCount int = 0
      oThisWeek.Find("a").Each(func(index int, sel *GoQuery.Selection) {
         iThisWeekCount++
      })
      sDoctor := v.Find(".index_top_in_name").Text()
      if 0 == iThisWeekCount {
          oThisWeek.Find("span").Each(func(index int, sel *GoQuery.Selection){

            sText := sel.Text()
            oFullSplit := strings.Split(sText, "|")
            oMapNumbers[sDoctor + "$" + oFullSplit[0]] = "0"
            sThisFullText += sText + "#"


              oThisWeekInfo = append(oThisWeekInfo, map[string]string{
                  "WeekInfo" : oFullSplit[0][:len(oFullSplit[0])-5],
                  "DateInfo" : oFullSplit[0][len(oFullSplit[0])-5:],
                  "NumberCount" : "0",
              })
          })
          if "" != sThisFullText{
            sThisFullText = sThisFullText[:len(sThisFullText) - 1]
          }
      }else{
          oThisWeek.Find("a").Each(func(index int, sel *GoQuery.Selection){
            sText := sel.Text()
            oNoFullSplit := strings.Split(sText, "|")
            oMapNumbers[sDoctor + "$" + oNoFullSplit[0]] = oNoFullSplit[1]
            sThisNoFullText += sText + "#"


            oThisWeekInfo = append(oThisWeekInfo, map[string]string{
                "WeekInfo" : oNoFullSplit[0][:len(oNoFullSplit[0])-5],
                "DateInfo" : oNoFullSplit[0][len(oNoFullSplit[0])-5:],
                "NumberCount" : oNoFullSplit[1],
            })
          })
          if "" != sThisNoFullText {
            sThisNoFullText = sThisNoFullText[:len(sThisNoFullText) - 1]
          }
      }
      //下周
      var sNextFullText string
      var sNextNoFullText string
      var iNextWeekCount int = 0
      oNextWeek.Find("a").Each(func(index int, sel *GoQuery.Selection){
          iNextWeekCount++
      })

      if 0 == iNextWeekCount {
          oNextWeek.Find("span").Each(func(index int, sel *GoQuery.Selection){
              sText := sel.Text()
              oFullSplit := strings.Split(sText, "|")
              oMapNumbers[sDoctor + "$" + oFullSplit[0]] = "0"
              sNextFullText += sText + "#"


              oNextWeekInfo = append(oNextWeekInfo, map[string]string{
                  "WeekInfo" : oFullSplit[0][:len(oFullSplit[0])-5],
                  "DateInfo" : oFullSplit[0][len(oFullSplit[0])-5:],
                  "NumberCount" : "0",
              })
          })
          if "" != sNextFullText {
              sNextFullText = sNextFullText[:len(sNextFullText) - 1]
          }

      }else{
          oNextWeek.Find("a").Each(func(index int, sel *GoQuery.Selection){
              sText := sel.Text()
              oNoFullSplit := strings.Split(sText, "|")
              oMapNumbers[sDoctor + "$" + oNoFullSplit[0]] = oNoFullSplit[1]
              sNextNoFullText += sText + "#"

              oNextWeekInfo = append(oNextWeekInfo, map[string]string{
                  "WeekInfo" : oNoFullSplit[0][:len(oNoFullSplit[0])-5],
                  "DateInfo" : oNoFullSplit[0][len(oNoFullSplit[0])-5:],
                  "NumberCount" : oNoFullSplit[1],
              })
          })
          if "" != sNextNoFullText {
              sNextNoFullText = sNextNoFullText[:len(sNextNoFullText) - 1]
          }

      }

      //iThisWeekCount = v.Find(".div_index_bottom .div_bottom_isweek .isweek_ind span").Text()
      sDept := v.Find(".index_top_in_exp").Text()
      oArrayDepts = append(oArrayDepts, sDept)
      oArrayDoctors = append(oArrayDoctors, sDoctor)

      //doctor struct save data
      oDoctors = append(oDoctors, &fiveHospistal.Doctor{
          Name : sDoctor,
          DeptName : sDept,
          ThisWeekInfo : oThisWeekInfo,
          NextWeekInfo : oNextWeekInfo,
      })

      fmt.Println(
        sDoctor +
        "("+ sDept +")" +
        "[这周:" +
        sThisFullText + "  " + sThisNoFullText +
        "][下周:" +
        sNextFullText + "  " + sNextNoFullText + "]")
        //fmt.Println(v.Find("index_top_in_name").Text())
        fmt.Println("=============================================")
    })

    //page loop
    oAllDom.Find(".list_paging").Each(func(_ int, v *GoQuery.Selection){
        iPage, err := strconv.Atoi(sPage)
        if nil != err {
            Error(err.Error())
            return
        }

        iPageNext := iPage + 1

        v.Find(".list_paging_btn").Each(func(index int, sel *GoQuery.Selection){
            if 2 == index {
                sHref, _ := sel.Attr("href")
                if "" == sHref {
                  //save to xlsx
                  //save to []string
                  var oDataArray [][]string

                  for k, v := range oArrayDoctors {

                      var sThisWeek string
                      var sNextWeek string
                      for kk, vv := range oMapNumbers {
                          if strings.Contains(kk, v + "$") {

                              oKV := strings.Split(kk, "$")
                              if len(oKV) < 2 {
                                  continue
                              }
                              if strings.Contains(kk, "本周"){
                                sThisWeek += oKV[1] + "[" + vv + "];"
                              }
                              if strings.Contains(kk, "下周"){
                                sNextWeek += oKV[1] + "[" + vv + "];"
                              }
                          }
                      }

                      oDataArray = append(oDataArray, []string{
                          oArrayDepts[k],
                          v,
                          sThisWeek,
                          sNextWeek,
                      })
                  }


                  //save to pgsql
                  //saveDataToPGSQL(oDataArray)
                  for _, v := range oDoctors{
                      v.SaveDataToPGSQL()
                  }



                  //before the save xls return
                  //the file append is too big
                  return

                  libs.SaveXlsx(
                    "第五医院信息表.xlsx",
                    "医生号源表",
                    []string{
                      "科室", "姓名", "这周", "下周",
                    },oDataArray)

                    return
                }
                //fmt.Println(strconv.Itoa(iPageNext))
                //fmt.Println("dohere")
                loopPageQuery(strconv.Itoa(iPageNext))
            }
        })
    })

}

//listen website with http
func ListenWebSite(){

    //run and query data

    // loopPageQuery("1")
    // fmt.Println(oDoctors)
    // return
    for{
      loopPageQuery("1")

      fmt.Print(">>>>>>>>>>>>>>>>>>>>> wait for : ")
      fmt.Println(oIni.QueryGap)
      time.Sleep(oIni.QueryGap)
    }
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
