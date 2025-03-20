package dags

import DAGModel "github.com/gotoflow/core/entities/dag"

type DagController struct {
	service *DagService
}

func ControllerConstructor () *DagController {
	return &DagController{
		service: ServiceConstructor(),
	}
}

func (dc *DagController) GetDags() ([]*DAGModel.DAG, error) {
	dag, err := dc.service.GetAll()
	return dag, err
}

func (dc *DagController) GetDag() {
	
}

func (dc *DagController) CreateDag(dag DagDTO) error {
	err := dc.service.CreateDag(dag)
	return err
	
}

func (dc *DagController) UpdateDag() {
	
}

