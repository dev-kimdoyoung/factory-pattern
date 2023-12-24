package main

type InterfaceVirtualMachine interface {
	getCpuCore() int
	getMemory() int
	getSSD() int
	getHDD() int
	getVendor() string
	setCpuCore(int)
	setMemory(int)
	setSSD(int)
	setHDD(int)
	setVendor(string)
}

type VirtualMachine struct {
	Core   int
	Memory int
	Ssd    int
	Hdd    int
	Vendor string
}

func (v *VirtualMachine) getCpuCore() int {
	return v.Core
}

func (v *VirtualMachine) setCpuCore(core int) {
	v.Core = core
}

func (v *VirtualMachine) getMemory() int {
	return v.Memory
}

func (v *VirtualMachine) setMemory(memory int) {
	v.Memory = memory
}

func (v *VirtualMachine) getSSD() int {
	return v.Ssd
}

func (v *VirtualMachine) setSSD(ssd int) {
	v.Ssd = ssd
}

func (v *VirtualMachine) getHDD() int {
	return v.Hdd
}

func (v *VirtualMachine) setHDD(hdd int) {
	v.Hdd = hdd
}

func (v *VirtualMachine) getVendor() string {
	return v.Vendor
}

func (v *VirtualMachine) setVendor(vendor string) {
	v.Vendor = vendor
}
