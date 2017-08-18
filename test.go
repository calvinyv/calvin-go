package main

// See https://docs.qingcloud.com/api/lb/index.html
// qingcloud loadBalancer and instance have a default strict Security Group(firewall),
// its only allow SSH and PING. So, first of all, you shoud manually add correct rules
// for all nodes and loadBalancers. You can simply add a rule by pass all tcp port traffic.
// The loadBalancers also need at least one EIP before create it, please allocate some EIPs,
// and set them in service annotation ServiceAnnotationLoadBalancerEipIds.

import (
	//	"github.com/golang/glog"
	"fmt"

	qcservice "github.com/yunify/qingcloud-sdk-go/service"
)

func main() {
	//	glog.V(3).Infof("test client %s", qcclient.InstanceStatusCeased)
	//fmt.Println("test client", qcclient.InstanceStatusCeased)
	status := []*string{qcservice.String("pending"), qcservice.String("active"), qcservice.String("stopped")}
	var name = "lb-lmld6diw"

	output, err := qcservice.LoadBalancerService.DescribeLoadBalancers(&qcservice.DescribeLoadBalancersInput{
		Status:     status,
		SearchWord: &name,
	})
	if err != nil {
		fmt.Print(err)
	}
	if len(output.LoadBalancerSet) == 0 {
		fmt.Print("no data")
	}
	for _, lb := range output.LoadBalancerSet {
		if lb.LoadBalancerName != nil && *lb.LoadBalancerName == "lb-lmld6diw" {
			fmt.Print(lb)
		}
	}

}
