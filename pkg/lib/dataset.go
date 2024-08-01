package easyenv

// constructor

func NewDataSet(keyName, value string) *DataSet {
	DataSet := new(DataSet)
	DataSet.keyName = keyName
	DataSet.SetValue(value)

	return DataSet
}

// setters

func (dtaset *DataSet) SetValue(value string) {
	dtaset.value = value
}

func (dtaset *DataSet) Remove() {
	dtaset.deleted = true
}

// getters

func (dtaset *DataSet) GetKey() string {
	return dtaset.keyName
}

func (dtaset *DataSet) GetValue() string {
	return dtaset.value
}
