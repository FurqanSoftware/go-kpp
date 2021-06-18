package kpp

import (
	"github.com/goccy/go-yaml"
)

type Metadata struct {
	Name            Names
	Type            Type
	Author          Strings
	Source          string
	SourceURL       string `yaml:"source_url"`
	License         string
	RightsOwner     string `yaml:"rights_owner"`
	Limits          Limits
	Validation      Validation
	ValidationFlags string  `yaml:"validator_flags"`
	Grading         Grading // DEPRECATED
	Scoring         Scoring
	Keywords        Strings
	UUID            string
	Libraries       Strings
	Languages       Strings
}

type Names map[string]string

func (n *Names) UnmarshalYAML(raw []byte) error {
	m := map[string]string{}
	err := yaml.Unmarshal(raw, &m)
	if err == nil {
		*n = m
		return nil
	}

	var v string
	err = yaml.Unmarshal(raw, &v)
	if err != nil {
		return err
	}
	*n = Names{"en": v}
	return nil
}

type Type string

const (
	TypePassFail = "pass-fail"
	TypeScoring  = "scoring"
)

type Strings []string

func (s *Strings) UnmarshalYAML(raw []byte) error {
	m := []string{}
	err := yaml.Unmarshal(raw, &m)
	if err == nil {
		*s = m
		return nil
	}

	var v string
	err = yaml.Unmarshal(raw, &v)
	if err != nil {
		return err
	}
	*s = Strings{v}
	return nil
}

type Limits struct {
	TimeMultiplier    float64 `yaml:"time_multiplier"`
	TimeSafetyMargin  float64 `yaml:"time_safety_margin"`
	Memory            int
	Output            int
	Code              int
	CompilationTime   int `yaml:"compilation_time"`
	CompilationMemory int `yaml:"compilation_memory"`
	ValidationTime    int `yaml:"validation_time"`
	ValidationMemory  int `yaml:"validation_memory"`
	ValidationOutput  int `yaml:"validation_output"`
}

type Validation string

const (
	ValidationDefault Validation = "default"
	ValidationCustom  Validation = "custom"
)

type Scoring struct {
	Objective string
}

type ScoringObjective string

const (
	ScoringObjectiveMin ScoringObjective = "min"
	ScoringObjectiveMax ScoringObjective = "max"
)

type Grading = Scoring
