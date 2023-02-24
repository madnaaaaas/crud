package logger

import "go.uber.org/zap"

func NewLogger(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	atom := zap.NewAtomicLevel()
	if err := atom.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	cfg.Level = atom
	cfg.InitialFields = map[string]interface{}{
		"service": "crud",
	}

	return cfg.Build()
}
