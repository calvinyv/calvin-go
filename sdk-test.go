package main

import (
	"fmt"
	"github.com/yunify/qingcloud-sdk-go/config"
	qs "github.com/yunify/qingcloud-sdk-go/service"
)

func main() {
	config, _ := config.New("FDAJUMZUJTGAAYQDQDNT", "zBqNi7zaQpKnHRruefrsE9S0jqFrT0OUVRvbcqNu")
	config.LogLevel = "debug"
	config.Zone = "pek3a"
	qcservice, _ := qs.Init(config)
	req := qs.CeaseInstancesInput{Instances: qs.StringSlice([]string{"i-iou9wqfy"})}
	instsvc, _ := qcservice.Instance(config.Zone)
	_, err := instsvc.CeaseInstances(&req)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//fmt.Printf("%s --- %d", *resp.Message, *resp.RetCode,resp.)

	//	i-sabu7kpr

}
