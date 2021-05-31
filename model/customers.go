package model

import (
	"encoding/json"
	"io"
)

type DbStore struct {
	List []DbEntity `json:"list"`
}

type DbEntity struct {
	CustomerIP string  `json:"ip"`
	TimeStamps []int64 `json:"data"`
}

func (dbStore *DbStore) ReadJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(dbStore)
	return err
}

func (dbStore *DbStore) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w).Encode(dbStore)
	return err
}

func ToDbStore(customers map[string][]int64) DbStore {

	var dbEntities []DbEntity
	for ip, data := range customers {
		entity := DbEntity{CustomerIP: ip, TimeStamps: data}
		dbEntities = append(dbEntities, entity)
	}
	return DbStore{List: dbEntities}
}

func ToCustomers(dbStore *DbStore) map[string][]int64 {

	customers := make(map[string][]int64)
	for _,dbEntity := range dbStore.List {
		customers[dbEntity.CustomerIP] = dbEntity.TimeStamps
	}

	return customers
}

