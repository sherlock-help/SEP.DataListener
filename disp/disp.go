package disp

import(
    SEP "SEP.DataListener/bizs"
)

func init(){

    //disp here
    //run web listener
    go SEP.ListenWebSite()

}
