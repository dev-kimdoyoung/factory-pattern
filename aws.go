package main

type Aws struct {
}

type Ec2 struct {
	VirtualMachine
}

type Rds struct {
	RelationDataBase
}

func (a *Aws) MakeVm() InterfaceVirtualMachine {
	return &Ec2{
		VirtualMachine{
			Core:   4,
			Memory: 8,
			Ssd:    200,
			Hdd:    200,
			Vendor: "aws",
		}}
}

func (a *Aws) MakeRdb() InterfaceRelationDataBase {
	return &Rds{
		RelationDataBase{
			DbName:     "rds",
			TableName:  []string{"a", "b", "c"},
			TableCount: 3,
		}}
}
