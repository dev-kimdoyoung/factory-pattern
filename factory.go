package main

import (
	"fmt"
	"strings"
)

type CloudServiceProviderFactory interface {
	MakeRdb() InterfaceRelationDataBase
	MakeVm() InterfaceVirtualMachine
}

func GetCloudServiceProviderFactory(provider string) (CloudServiceProviderFactory, error) {
	switch strings.ToUpper(provider) {
	case "AWS":
		return &Aws{}, nil
	case "GCP":
		return &Gcp{}, nil
	}
	return nil, fmt.Errorf("Wrong Cloud Service Provider type passed")
}
