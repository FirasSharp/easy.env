package easyenv

import (
	"errors"
	"fmt"
	"sync"
)

func createTables(connection *Connection) error {
	db := connection.db
	_, err := db.Exec("CREATE TABLE projects(projectID TEXT PRIMARY KEY, projectName TEXT, path TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE templates(templateID TEXT PRIMARY KEY, templateName TEXT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE templateValues(keyName TEXT PRIMARY KEY, templateID TEXT, value TEXT, FOREIGN KEY(templateID) REFERENCES templates(templateID))")
	if err != nil {
		return err
	}
	return nil
}

func saveDataInDB(connection *Connection) error {
	var err error
	var errorText string
	wg := new(sync.WaitGroup)

	var projectError error
	var templateError error

	wg.Add(2)

	go saveProjects(connection, &projectError, wg)
	go saveTemplates(connection, &templateError, wg)

	wg.Wait()

	var templateEnvError error

	wg.Add(1)

	go saveEnvTemplates(connection, &templateEnvError, wg)

	wg.Wait()

	if projectError != nil {
		errorText = fmt.Sprintf("An error occurred while saving the project. Details: %s\n", projectError.Error())
	}

	if templateError != nil {
		errorText = fmt.Sprintf("%sAn error occurred while saving the templates. Details: %s\n", errorText, templateError.Error())
	}

	if templateEnvError != nil {
		errorText = fmt.Sprintf("%sAn error occurred while saving the env in templates. details: %s\n", errorText, templateEnvError.Error())
	}

	if len(errorText) > 0 {
		err = errors.New(errorText)
	}

	return err
}

func selectProjects(connection *Connection) (map[string]*Project, error) {
	result := make(map[string]*Project)
	db := connection.db
	query := "SELECT * FROM projects"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var projectID string
		var projectName string
		var path string
		err := rows.Scan(projectID, projectName, path)
		if err != nil {
			return nil, err
		}

		project := NewProject(projectName, path)
		project.projectID = projectID

		result[project.projectID] = project
	}

	return result, nil
}

func selectTemplates(connection *Connection) (map[string]*Template, error) {
	result := make(map[string]*Template)
	db := connection.db
	query := "SELECT * FROM templates"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var templateID string
		var templateName string

		err := rows.Scan(&templateID, &templateName)
		if err != nil {
			return nil, err
		}

		template := NewTemplate(templateName)
		template.templateID = templateID

		template.values, err = selectTemplateEnvironments(connection, template.templateID)
		if err != nil {
			return nil, err
		}

		result[template.templateID] = template
	}

	return result, nil
}

func selectTemplateEnvironments(connection *Connection, templateID string) ([]*DataSet, error) {
	var result []*DataSet
	db := connection.db
	query := "SELECT keyName, value FROM templateValues WHERE templateID = ?"

	rows, err := db.Query(query, templateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var keyName string
		var value string
		err := rows.Scan(&keyName, &value)
		if err != nil {
			return nil, err
		}

		env := NewDataSet(keyName, value)
		result = append(result, env)
	}

	return result, nil
}

func saveProjects(connection *Connection, errorResult *error, wg *sync.WaitGroup) {
	defer wg.Done()
	db := connection.db

	tx, err := db.Begin()
	if err != nil {
		*errorResult = err
		return
	}

	for _, project := range connection.projects {
		if project.deleted {
			query := "DELETE FROM projects WHERE projectID = ?"
			_, err := tx.Exec(query, project.GetProjectID())
			if err != nil {
				tx.Rollback()
				*errorResult = err
				return
			}
			continue
		}

		query := "INSERT INTO projects(projectID, projectName, path) VALUES(?, ?, ?) ON CONFLICT(projectID) DO UPDATE SET projectName = ?"
		_, err := tx.Exec(query, project.GetProjectID(), project.GetProjectName(), project.GetPath(), project.GetProjectName())
		if err != nil {
			tx.Rollback()
			*errorResult = err
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		*errorResult = err
	}
}

func saveTemplates(connection *Connection, errorResult *error, wg *sync.WaitGroup) {
	defer wg.Done()
	db := connection.db

	tx, err := db.Begin()
	if err != nil {
		*errorResult = err
		return
	}

	for _, template := range connection.templates {
		if template.deleted {
			query := "DELETE FROM templates WHERE templateID = ?"
			_, err := tx.Exec(query, template.GetTemplateID())
			if err != nil {
				tx.Rollback()
				*errorResult = err
				return
			}
			continue
		}

		query := "INSERT INTO templates(templateID, templateName) VALUES(?, ?) ON CONFLICT(templateID) DO UPDATE SET templateName = ?"
		_, err := tx.Exec(query, template.GetTemplateID(), template.GetTemplateName(), template.GetTemplateName())
		if err != nil {
			tx.Rollback()
			*errorResult = err
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		*errorResult = err
	}
}

func saveEnvTemplates(connection *Connection, errorResult *error, wg *sync.WaitGroup) {
	defer wg.Done()
	db := connection.db

	tx, err := db.Begin()
	if err != nil {
		*errorResult = err
		return
	}

	for _, template := range connection.templates {
		for _, templateEnv := range template.values {
			if templateEnv.deleted {
				query := "DELETE FROM templateValues WHERE keyName = ? AND templateID = ?"
				_, err := tx.Exec(query, templateEnv.GetKey(), template.GetTemplateID())
				if err != nil {
					tx.Rollback()
					*errorResult = err
					return
				}
				continue
			}

			_, err := tx.Exec("REPLACE INTO templateValues(keyName, templateID, value) VALUES(?, ?, ?)", templateEnv.GetKey(), template.GetTemplateID(), templateEnv.GetValue())
			if err != nil {
				tx.Rollback()
				*errorResult = err
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		*errorResult = err
	}
}
