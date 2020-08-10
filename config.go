package main

// ApplicationConfig - the configuration required for the application to run
type ApplicationConfig struct {
	AwsAccessKeyID     string   `validate:"required" mapstructure:"access_key_id"`
	AwsSecretAccessKey string   `validate:"required" mapstructure:"secret_access_key"`
	AwsS3Bucket        string   `validate:"required" mapstructure:"bucket"`
	AwsRegion          string   `validate:"required" mapstructure:"region"`
	DatabaseDriver     string   `validate:"required" mapstructure:"driver"`
	DatabaseUsername   string   `validate:"required" mapstructure:"username"`
	DatabasePassword   string   `validate:"required" mapstructure:"password"`
	DatabaseHost       []string `validate:"required" mapstructure:"host"`
	Database           string   `validate:"required" mapstructure:"database"`
	DatabaseCollection string   `validate:"required" mapstructure:"collection"`
}
