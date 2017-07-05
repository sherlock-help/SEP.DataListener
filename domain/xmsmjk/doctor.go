package xmsmjk

import(

    //bulit-in
  //  "fmt"
    "time"

    //extends
    "SEP.DataListener/libs/dao"

    . "bakerstreet-club/logs"

)

var oPgsql *dao.PGSql

func init(){
    oPgsql = dao.GetPgObj()
}

//the doctor in xmsmjk
type Doctor struct {

    //the doctor's name
    Name string
    //what the dept doctor in
    DeptName string
    //this week
    ThisWeekInfo []map[string]string
    //next week
    NextWeekInfo []map[string]string
}
/*
  data demo:
  {
  	医生：{
  		"科室" : {

  		},
  		"这周" : [{
  			"星期" : "",
  			"日期" : "",
  			"号源" : ""
  		}],
  		"下周" : []
  	}
  }
*/

//save to postgresql
func (oDoctor *Doctor) SaveDataToPGSQL(){

    //get the data from doctor

    //this week
    oThisWeek := oDoctor.ThisWeekInfo

    for _, v := range oThisWeek {

      //get from dictionary
      var sWeekInfo string
      if _, ok := v["WeekInfo"]; ok {
          sWeekInfo = v["WeekInfo"]
      }
      var sDateInfo string
      if _, ok := v["DateInfo"]; ok {
          sDateInfo = v["DateInfo"]
      }
      var sNumberCount string
      if _, ok := v["NumberCount"]; ok {
          sNumberCount = v["NumberCount"]
      }

      // fmt.Println(map[string]string{
      //     "DoctorName" : oDoctor.Name,
      //     "DeptName" : oDoctor.DeptName,
      //     "WeekType" : "下周",
      //     "WeekInfo" : sWeekInfo,
      //     "DateInfo" : sDateInfo,
      //     "NumberCount" : sNumberCount,
      // })

      oDoctor.inOrUpNumber(map[string]string{
          "DoctorName" : oDoctor.Name,
          "DeptName" : oDoctor.DeptName,
          "WeekType" : "本周",
          "WeekInfo" : sWeekInfo,
          "DateInfo" : sDateInfo,
          "NumberCount" : sNumberCount,
      })
    }

    //next week
    oNextWeek := oDoctor.NextWeekInfo
    for _, v := range oNextWeek {

      //get from dictionary
      var sWeekInfo string
      if _, ok := v["WeekInfo"]; ok {
          sWeekInfo = v["WeekInfo"]
      }
      var sDateInfo string
      if _, ok := v["DateInfo"]; ok {
          sDateInfo = v["DateInfo"]
      }
      var sNumberCount string
      if _, ok := v["NumberCount"]; ok {
          sNumberCount = v["NumberCount"]
      }

      oDoctor.inOrUpNumber(map[string]string{
          "DoctorName" : oDoctor.Name,
          "DeptName" : oDoctor.DeptName,
          "WeekType" : "下周",
          "WeekInfo" : sWeekInfo,
          "DateInfo" : sDateInfo,
          "NumberCount" : sNumberCount,
      })

    }

}


//insert or update number
func (oDoctor *Doctor) inOrUpNumber(oMap map[string]string){

      //check first
      if nil == oMap {
            Error("sorry, the map is empty, so not save to database")
            return
      }


      //give hour here in map
      sNowHour := time.Now().Format("15")
      oMap["DayHour"] = sNowHour
      // fmt.Println(iNowHour)
      // return

      //query first and think insert or update
      oMapQuery := map[string] dao.PGQueryObj {
    		"five_number" : dao.PGQueryObj{
    			Fields : []string{
              "DayHour",
          },
          Orders : []string{
          },
    			ObjWhere : []dao.PGQueryWhere {
            *&dao.PGQueryWhere{
              Field : "DeptName",
    					Op : "equal",
    					Val : oMap["DeptName"],
            },
            *&dao.PGQueryWhere{
              Field : "DoctorName",
    					Op : "equal",
    					Val : oMap["DoctorName"],
            },
            *&dao.PGQueryWhere{
              Field : "DateInfo",
    					Op : "equal",
    					Val : oMap["DateInfo"],
            },
            *&dao.PGQueryWhere{
              Field : "DayHour",
    					Op : "equal",
    					Val : oMap["DayHour"],
            },
    			},
    		},
  	}

      if len(oPgsql.QueryData(oMapQuery)["five_number"]) > 0 {
        oPgsql.Update("five_number",
          oMap,
          map[string]string{
            "DeptName" : oMap["DeptName"],
            "DoctorName" : oMap["DoctorName"],
            "DateInfo" : oMap["DateInfo"],
            "DayHour" : oMap["DayHour"],
          })
      }else{
        oPgsql.Insert("five_number", oMap)
      }
}
