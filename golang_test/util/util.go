package main

import (
	"fmt"
	"strings"
)

func main() {

	// split 호출
	s := "pg=1&s=TO&a=DITO&dc=1DV&e2=개 종류&r=10&rc=23&d=2FaG4NMgw4iEQeCgTX&p=1&qs=q%3D%25EA%25B7%2580%25EC%2597%25AC%25EC%259A%25B4%2520%25EA%25B0%2595%25EC%2595%2584%25EC%25A7%2580%26w%3Dtot%26nil_profile%3Dsuggest%26DA%3DGIQ%26rq%3D%26sq%3D%25EA%25B7%2580%25EC%2597%25AC%25EC%259A%25B4%2520%25EA%25B0%2595%25E3%2585%2587%25EC%25A7%2580%26o%3D1%26sugo%3D10&userip=211.114.121.245&ud=MSNv1pXyj91B_180103155543723&uk=Wkx@bsuFo2QAAjRssq8AAACi&mk=Wkx@i9PnaTQAABjHZlAAAACJ&uc=2&xy=516_243/1366_650&ct=9&io=ki&b=&w=tot&q=귀여운 강아지&ext=encytype%3DCAT20%7C%7CNORM__DeLiM__bucketType%3D__DeLiM__dl%3DNSJ20_1DV23_IIM20_KTP1_TWA3_BRC3_VO25_TWB7_NNS3_KAS3_KOQ10_0BL5_DCC20_1TH10_EOA10_MBK10_DRU10_RME10_VLY10_LEE10_MCO10_CLO6&u=https%3A%2F%2Fm.search.daum.net%2Fkakao%3Fw%3Dtot%26DA%3D1DV%26rtmaxcoll%3D1DV%26ency%3Dsq%253D%25EA%25B0%259C%2520%25EC%25A2%2585%25EB%25A5%2598%26q%3D%25EB%25B9%2584%25EC%2588%2591%2520%25ED%2594%2584%25EB%25A6%25AC%25EC%25A0%259C&callback=SF.M.smartLogCallBack&_=1514962571503"
	sp := "&"
	split(s, sp)

}

func split(s string, sp string) {
	st := strings.Split(s, sp)
	for _, a := range st {
		fmt.Println(a)
	}
}
