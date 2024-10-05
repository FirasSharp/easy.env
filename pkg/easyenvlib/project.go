package easyenv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type Project struct {
	projectID   string
	projectName string
	path        string // absolute path of the project containing the .env file
	deleted     bool
	values      map[string]*DataSet
}

// constructor

func NewProject(projectName, path string) *Project {
	project := new(Project)

	project.projectID = uuid.NewString()
	project.SetProjectName(projectName)
	project.SetPath(path)
	project.values = make(map[string]*DataSet)

	return project
}

// setters

func (prj *Project) SetProjectName(value string) {
	prj.projectName = value
}

func (prj *Project) SetPath(value string) {
	prj.path = filepath.Join(value, ".env")
}

func (prj *Project) AddEnvironment(keyName, value string) (*DataSet, error) {

	_, ok := prj.GetEnvironmentByKey(keyName)

	if ok == nil {
		return nil, fmt.Errorf("an enviorment with the key %s already exists", keyName)
	}

	env := NewDataSet(keyName, value)
	prj.values[keyName] = env
	return env, nil
}

func (prj *Project) Remove() {
	prj.deleted = true
}

func (prj *Project) RemoveEnviorment(keyName string) {
	delete(prj.values, keyName)
}

func (prj *Project) RemoveAllEnviorments() error {

	prj.values = make(map[string]*DataSet)

	return nil
}

// getters

func (prj *Project) GetProjectName() string {
	return prj.projectName
}

func (prj *Project) GetProjectID() string {
	return prj.projectID
}

func (prj *Project) GetPath() string {
	return prj.path
}

func (prj *Project) GetEnvironments() map[string]*DataSet {
	return prj.values
}

func (prj *Project) GetEnvironmentByKey(keyName string) (*DataSet, error) {
	env, ok := prj.values[keyName]

	if !ok {
		return nil, fmt.Errorf("no enviorment found with the key %s", keyName)
	}

	return env, nil
}

func (prj *Project) LoadEnvironmentsFromFile() error {
	envPath := prj.GetPath()
	data, err := os.ReadFile(envPath)

	if err != nil {
		return err
	}

	enviorments := string(data)

	enviorments = strings.ReplaceAll(enviorments, "\r", "")

	envs := strings.Split(enviorments, "\n")

	for _, env := range envs {

		if len(env) == 0 { // in case the file has an empty line
			continue
		}

		splittedEnv := strings.Split(env, "=")
		prj.AddEnvironment(splittedEnv[0], splittedEnv[1])
	}

	return nil
}

// functionalities

func (prj *Project) SaveEnvironmentsToFile() error {
	var err error

	envPath := prj.GetPath()
	os.Remove(envPath)

	envString := createEnvString(prj.values)

	err = os.WriteFile(envPath, []byte(envString), 0644)

	if err != nil {
		return err
	}
	return err
}

func createEnvString(environments map[string]*DataSet) string {
	var result string

	for _, env := range environments {
		result += fmt.Sprintf("%s=%s\n", env.keyName, env.value)
	}

	return result
}
