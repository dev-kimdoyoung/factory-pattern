package main

type RelationDataBase struct {
	DbName     string
	TableName  []string
	TableCount int
}
type InterfaceRelationDataBase interface {
	GetDbName() string
	UpdateDbName(string)
	GetTableNameList() []string
	GetTableCount() int
}

func (rdb *RelationDataBase) GetDbName() string {
	return rdb.DbName
}

func (rdb *RelationDataBase) UpdateDbName(name string) {
	rdb.DbName = name
}

func (rdb *RelationDataBase) GetTableNameList() []string {
	return rdb.TableName
}

func (rdb *RelationDataBase) GetTableCount() int {
	return rdb.TableCount
}
