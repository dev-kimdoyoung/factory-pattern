package main

import (
	"fmt"
)

func main() {
	awsFactory, _ := GetCloudServiceProviderFactory("aws")
	gcpFactory, _ := GetCloudServiceProviderFactory("gcp")

	ec2 := awsFactory.MakeVm()
	rds := awsFactory.MakeRdb()

	vm := gcpFactory.MakeVm()
	rdb := gcpFactory.MakeRdb()

	fmt.Println("aws ec2 vendor : ", ec2.getVendor())
	fmt.Println("aws rds db name : ", rds.GetDbName())

	fmt.Println("gcp vm vendor : ", vm.getVendor())
	fmt.Println("gcp rdb db name : ", rdb.GetDbName())
}
