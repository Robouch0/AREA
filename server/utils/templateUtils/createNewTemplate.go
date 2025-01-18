//
// EPITECH PROJECT, 2025
// AREA
// File description:
// createNewTemplate
//

package template_utils

import (
	"html/template"
)

func CreateNewTemplate(name, text string) (*template.Template, error) {
	t := template.New(name)
	t, err := t.Parse(text)
	if err != nil {
		return nil, err
	}
	return t, nil
}
