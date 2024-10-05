package easyenv

import (
	"fmt"

	"github.com/google/uuid"
)

type Template struct {
	templateID   string
	templateName string
	deleted      bool
	values       map[string]*DataSet
}

// Constructor
func NewTemplate(templateName string) *Template {
	template := new(Template)
	template.templateID = uuid.NewString()
	template.values = make(map[string]*DataSet)
	template.SetTemplateName(templateName)
	return template
}

// Getters
func (template *Template) GetTemplateID() string {
	return template.templateID
}

func (template *Template) GetTemplateName() string {
	return template.templateName
}

func (template *Template) GetEnvironments() map[string]*DataSet {
	return template.values
}

func (template *Template) GetEnvironmentByKey(keyName string) (*DataSet, error) {

	env, ok := template.values[keyName]

	if !ok {
		return nil, fmt.Errorf("no environment found with the key %s", keyName)
	}
	return env, nil
}

// Setters
func (template *Template) SetTemplateName(templateName string) {
	template.templateName = templateName
}

func (template *Template) AddEnvironment(keyName, value string) (*DataSet, error) {

	_, ok := template.GetEnvironmentByKey(keyName)

	if ok == nil {
		return nil, fmt.Errorf("an environment with the key %s already exists", keyName)
	}

	env := NewDataSet(keyName, value)
	template.values[keyName] = env
	return env, nil
}

func (template *Template) Remove() {
	template.deleted = true
	for _, data := range template.values {
		data.Remove()
	}
}

func (template *Template) RemoveEnvironment(keyName string) {
	delete(template.values, keyName)
}

// This method will remove the environment
func (template *Template) RemoveAllEnvironments() {
	template.values = make(map[string]*DataSet)
}
