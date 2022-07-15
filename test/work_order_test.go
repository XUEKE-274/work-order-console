package test

import (
	"github.com/imroc/req"
	"testing"
	"time"
)

func TestPageList(t *testing.T)  {




	uri := "https://mk-test2.dustess.com/java-work-service/api/v1/ticket_query/page-list"
	m := make(map[string]interface{})

	m["pageType"] = "ALL"
	m["sourceType"] = "PC_BACKSTAGE"
	m["orderBy"] = "updateTime"
	m["orderDirect"] = "desc"
	//m["ticketNumberOrTitle"] = time.Now().String()
	m["Index"] = 3
	m["pageSize"] = 20
	body := req.BodyJSON(&m)
	hs := make(map[string]string)



	hs["Authorization"] = "qw.b614a1be-00c6-11ed-be1f-bece3d10df62"
	h := req.HeaderFromStruct(hs)


	var total int64 = 0


	for i := 0; i < 20; i++{
		go func() {
			for ;;{


			total ++
			res ,_ :=req.Post(uri, body, h)
			println(res.ToString())
			println(total)


			}
		}()
	}
	time.Sleep(time.Hour)






}


