package dags

import (
	"fmt"
	"reflect"

	DAGModel "github.com/gotoflow/core/entities/dag"
	internal_database "github.com/gotoflow/gotoflow-api/src/handlers/database"
	"gorm.io/gorm"
)

type DagService struct {
	database *gorm.DB
}

func ServiceConstructor() *DagService {
	return &DagService{
		database: internal_database.GetClientDatabase(),
	}
}

func (ds DagService) CreateDag (dagModel DagDTO) error {
	dagInsert := &DAGModel.DAG{
		DagId: dagModel.DagId,
		Path: dagModel.DagPath,
	}
	err := ds.database.Create(dagInsert)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ds DagService) GetAll() ([]*DAGModel.DAG, error) {
	getStruct := []*DAGModel.DAG{}
	err := ds.database.Model(getStruct).Scan(&getStruct)
	if err.Error != nil {
		fmt.Println(reflect.TypeOf(err.Error))
	}
	// if err.Error != nil && err.Error == gorm.ErrRelat{
	// 	return nil, err
	return getStruct, err.Error
}