package easyenv

import (
	"fmt"

	"github.com/google/uuid"
)

// Constructor
func NewTemplate(templateName string) *Template {
	template := new(Template)
	template.templateID = uuid.NewString()
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

func (template *Template) GetEnvironments() []*DataSet {
	return template.values
}

func (template *Template) GetEnvironmentByKey(keyName string) (*DataSet, error) {
	for _, env := range template.values {
		if env.keyName == keyName {
			return env, nil
		}
	}
	return nil, fmt.Errorf("no environment found with the key %s", keyName)
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
	template.values = append(template.values, env)
	return env, nil
}

func (template *Template) Remove() {
	template.deleted = true
	for _, data := range template.values {
		data.Remove()
	}
}

func (template *Template) RemoveEnvironment(keyName string) {
	var tmp []*DataSet
	for _, env := range template.values {
		if env.keyName != keyName {
			tmp = append(tmp, env)
		}
	}
	template.values = tmp
}

// This method will remove the environment
func (template *Template) RemoveAllEnvironments() {
	template.values = []*DataSet{}
}
