package initialize

import (
	"fmt"
	"github.com/black-dragon74/rht-trn/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type DataStore struct {
	Students []types.Student
	lgr      *zap.Logger
}

func NewDataStore(lgr *zap.Logger) *DataStore {
	lgr.Info("[Initialize] [DataStore] [NewDataStore] Loading the data in memory")

	store := &DataStore{
		lgr: lgr,
	}
	store.Refresh()

	//	Otherwise, we have our list
	lgr.Info("[Initialize] [DataStore] [NewDataStore] Successfully loaded the data in memory")
	return store
}

func (d *DataStore) AddStudent(student *types.Student) {
	d.lgr.Info(fmt.Sprintf("[Initialize] [DataStore] [AddStudent] Adding student: %s", student.Name))

	d.Students = append(d.Students, *student)
	go d.writeToDisk()
}

func (d *DataStore) Refresh() {
	d.lgr.Info("[Initialize] [DataStore] Refreshing the store...")

	var students []types.Student

	data, err := os.ReadFile("students.yaml")
	if err != nil {
		d.lgr.Error("[Initialize] [DataStore] [Refresh] Error while reading the students.yaml file")
	}

	err = yaml.Unmarshal(data, &students)
	if err != nil {
		d.lgr.Error("[Initialize] [DataStore] [Refresh] Error while unmarshaling YAML to go types")
	}

	d.Students = students
}

func (d *DataStore) writeToDisk() {
	d.lgr.Info("[Initialize] [DataStore] [WriteToDisk] Attempting to flush contents to disk")

	//	Marshal the contents
	data, err := yaml.Marshal(d.Students)
	if err != nil {
		d.lgr.Error("[Initialize] [DataStore] [WriteToDisk] Error in marhsalling the YAML")
	}

	err = os.WriteFile("students.yaml", data, 0644)
	if err != nil {
		d.lgr.Error("[Initialize] [DataStore] [WriteToDisk] Error in flushing to disk")
	}
}
