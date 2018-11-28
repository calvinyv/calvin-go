//package main
//
//// See https://docs.qingcloud.com/api/lb/index.html
//// qingcloud loadBalancer and instance have a default strict Security Group(firewall),
//// its only allow SSH and PING. So, first of all, you shoud manually add correct rules
//// for all nodes and loadBalancers. You can simply add a rule by pass all tcp port traffic.
//// The loadBalancers also need at least one EIP before create it, please allocate some EIPs,
//// and set them in service annotation ServiceAnnotationLoadBalancerEipIds.
//
//import (
//	"fmt"
//	// "io/ioutil"
//	// "log"
//	"os"
//	//"path/filepath"
//	//"path/filepath"
//	"strconv"
//)
//
//func main() {
//	// files, err := ioutil.ReadDir("/dev/disk/by-id")
//	// if err != nil {
//	// 	log.Fatal(err)
//	// }
//
//	// lstr, err := os.Readlink("/bin/sh")
//	// if err != nil {
//	// 	os.Exit(1)
//	// }
//	//path := "data/var/lib/kubelet/plugins/kubernetes.io/flexvolume/qingcloud/flex-volume/mounts/pvc-ab52a0f1-fa62-11e7-806e-52541fed8204"
//
//	//fmt.Println(os.Args[1])
//
//
//
//	//if len(os.Args) <= 1 {
//	//	fmt.Println("Pls input line number")
//	//	return
//	//}
//	//line, err := strconv.Atoi(os.Args[1])
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//if line <= 0 {
//	//	fmt.Println("wrong line number")
//	//	return
//	//}
//	//
//	//printStarTree(line)
//
//
//
//
//	//fmt.Println(filepath.Base(path))
//
//}
//
//func printStarTree(line int) error {
//
//	start := "*"
//	for i := 0; i < line; i++ {
//		fmt.Println(start)
//		start = start + "**"
//	}
//	return nil
//}
package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/calvinyv/calvin-go/api/test"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterTestServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
