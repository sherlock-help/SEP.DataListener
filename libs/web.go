package libs

import(
    //bulit-in
    "net/http"
    "io/ioutil"
    "net/url"
    "strings"

    //extends
    . "bakerstreet-club/logs"
    GoQuery "github.com/bakerstreet-club/goquery"
//    iconv "github.com/sherlock-help/iconv-go"
)

//param:
//sURL
//oURLBring the first is param; and second is header
func QueryByURL(sURL string, oURLBring ...map[string]string) interface{}{

    //this go query
    return GoQueryByURL(sURL)

    //query by http
    return DoQueryByURL(sURL, oURLBring)
}

//go query
func GoQueryByURL(sURL string) interface{} {

    //check first
    if "" == sURL {
        Error("sorry, the param named sURL is can not be empty!")
        return ""
    }

    //go query begin
    doc, err := GoQuery.NewDocument(sURL)
    if nil != err {
        Error(err.Error())
    }
    var sReText string
    doc.Find("script").Each(func(i int, s *GoQuery.Selection){
        //for Each

        sReText = s.Text()
    })
    return sReText
}

func DoQueryByURL(sURL string, oURLBring []map[string]string) interface{} {

    //check first
    if sURL == "" {
        Error("sorry, the param named sURL is can not be emtpy! ")
        return nil
    }

    //query here

    oParam := url.Values{}
    if len(oURLBring) > 0 {
      //set param
       for k, v := range oURLBring[0] {
          oParam.Set(k, v)
       }
    }
    bodyPost := ioutil.NopCloser(strings.NewReader(oParam.Encode()))
    clientPost := &http.Client{}

    rq, _ := http.NewRequest("GET",
      sURL,
      bodyPost)
    //set header
    if len(oURLBring) > 1 {
      //set Header
      for k, v := range oURLBring[1] {
          rq.Header.Set(k, v)
      }
    }
   rq.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
   //rq.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
   rq.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
   rq.Header.Set("Cache-Control", "max-age=0")
   rq.Header.Set("Connection", "keep-alive")

   respPost, _ := clientPost.Do(rq)
   defer respPost.Body.Close()
   dataPost, _ := ioutil.ReadAll(respPost.Body)

   //change encode
   //out := make([]byte,len(dataPost))
  // iconv.Convert(dataPost, out, "gb2312", "utf-8")

   return string(dataPost)
}
