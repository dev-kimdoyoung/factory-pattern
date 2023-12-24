package main

type Gcp struct {
}

type Vm struct {
	VirtualMachine
}

type Rdb struct {
	RelationDataBase
}

func (g *Gcp) MakeVm() InterfaceVirtualMachine {
	return &Vm{
		VirtualMachine{
			Core:   4,
			Memory: 8,
			Ssd:    200,
			Hdd:    200,
			Vendor: "google",
		}}
}

func (g *Gcp) MakeRdb() InterfaceRelationDataBase {
	return &Rds{
		RelationDataBase{
			DbName:     "rdb",
			TableName:  []string{"a", "b", "c"},
			TableCount: 3,
		}}
}
