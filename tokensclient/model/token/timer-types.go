package token

import (
	"encoding/json"
	"strings"
)

type TimerDefinition struct {
	Duration    int                `yaml:"duration,omitempty" mapstructure:"duration,omitempty" json:"duration,omitempty"`
	Description string             `yaml:"description,omitempty" mapstructure:"description,omitempty" json:"description,omitempty"`
	Actions     []ActionDefinition `yaml:"actions,omitempty" mapstructure:"actions,omitempty" json:"actions,omitempty"`
}

// Timer the Timer type (and ActionDefinition) exists in three places: tpm-tokens-client, tpm-tokens and leas-cab-token-jobs
type Timer struct {
	Pkey            string           `yaml:"pkey,omitempty" mapstructure:"pkey,omitempty" json:"pkey,omitempty"`
	Id              string           `yaml:"id,omitempty" mapstructure:"id,omitempty" json:"id,omitempty"`
	TokenId         string           `yaml:"token-id,omitempty" mapstructure:"token-id,omitempty" json:"token-id,omitempty"`
	Expires         string           `yaml:"expires,omitempty" mapstructure:"expires,omitempty" json:"expires,omitempty"`
	Outdated        bool             `yaml:"outdated,omitempty" mapstructure:"outdated,omitempty" json:"outdated,omitempty"`
	TimerDefinition *TimerDefinition `yaml:"definition,omitempty" mapstructure:"definition,omitempty" json:"definition,omitempty"`
	TTL             int              `yaml:"ttl,omitempty" mapstructure:"ttl,omitempty" json:"ttl,omitempty"`
}

func (timer *Timer) MarkAsOutdated() {
	timer.TimerDefinition = nil
	timer.Outdated = true
}

func (timer *Timer) ToJSON() ([]byte, error) {
	return json.Marshal(timer)
}

func (timer *Timer) MustToJSON() []byte {
	b, err := json.Marshal(timer)
	if err != nil {
		panic(err)
	}

	return b
}

func WellFormTimerId(id string) string {
	return strings.ToUpper(id)
}

func DeserializeTimer(b []byte) (*Timer, error) {
	timer := Timer{}
	err := json.Unmarshal(b, &timer)
	if err != nil {
		return nil, err
	}

	return &timer, nil
}
