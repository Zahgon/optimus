package plugin

import (
	"github.com/AlecAivazis/survey/v2"
)

type Options struct {
	DryRun bool
}

// USED in Question Validations
type vFactory struct{}

func (*vFactory) NewFromRegex(re, message string) survey.Validator {
	_ = "STUB: not implemented"
	return *new(survey.Validator)
}

var ValidatorFactory = new(vFactory)

type Question struct {
	Name        string   `yaml:",omitempty"`
	Prompt      string   `yaml:",omitempty"`
	Help        string   `yaml:",omitempty"`
	Default     string   `yaml:",omitempty"`
	Multiselect []string `yaml:",omitempty"`

	SubQuestions []SubQuestion `yaml:",omitempty"`

	Regexp          string `yaml:",omitempty"`
	ValidationError string `yaml:",omitempty"`
	MinLength       int    `yaml:",omitempty"`
	MaxLength       int    `yaml:",omitempty"`
	Required        bool   `yaml:",omitempty"`
}

func (q *Question) IsValid(value string) error { _ = "STUB: not implemented"; return nil }

type SubQuestion struct {
	// IfValue is used as an if condition to match with user input
	// if user value matches this only then ask sub questions
	IfValue   string
	Questions Questions
}

type Questions []Question

func (q Questions) Get(name string) (Question, bool) {
	_ = "STUB: not implemented"
	return *new(Question), false
}

type Answer struct {
	Question Question
	Value    string
}

type Answers []Answer

func (ans Answers) Get(name string) (Answer, bool) {
	_ = "STUB: not implemented"
	return *new(Answer), false
}
