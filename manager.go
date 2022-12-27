package geetestbot

import (
	"errors"
	"fmt"
)

// Manager 管理器
type Manager struct {
	config *Config
	apis   map[string]Api

	enableLog bool
}

// Api 获取Api，如果不指定name，则读取配置的default指定的name
func (m *Manager) Api(names ...string) (Api, error) {
	var err error
	var name = m.config.Default
	if len(names) != 0 {
		name = names[0]
	}

	rtr, ok := m.apis[name]
	if ok {
		return rtr, nil
	}

	rtr, err = m.resolve(name)
	if err != nil {
		return nil, err
	}

	if m.apis == nil {
		m.apis = make(map[string]Api, 0)
	}

	m.apis[name] = rtr

	return m.apis[name], err
}

func (m *Manager) resolve(name string) (Api, error) {
	config, ok := m.config.Apis[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("api [%s] does not have a config.", name))
	}

	a, err := NewApiFromConfig(&config)
	if err != nil {
		return nil, err
	}

	return a, nil
}
