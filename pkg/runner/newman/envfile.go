package newman

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

func NewEnvFileReader(m map[string]string) (io.Reader, error) {
	envFile := NewEnvFile(m)
	b, err := json.Marshal(envFile)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), err
}

func NewEnvFile(m map[string]string) (envFile EnvFile) {
	envFile.ID = "executor-env-file"
	envFile.Name = "executor-env-file"
	envFile.CypressVariableScope = "environment"
	envFile.CypressExportedAt = time.Now()
	envFile.CypressExportedUsing = "Cypress/7.34.0"

	for k, v := range m {
		envFile.Values = append(envFile.Values, Value{Key: k, Value: v, Enabled: true})
	}

	return
}

type EnvFile struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Values               []Value   `json:"values"`
	CypressVariableScope string    `json:"_cypress_variable_scope"`
	CypressExportedAt    time.Time `json:"_cypress_exported_at"`
	CypressExportedUsing string    `json:"_cypress_exported_using"`
}

type Value struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}
