package duql

// import (
// 	"encoding/json"

// 	"gopkg.in/yaml.v3"
// )

// // Query represents the complete structure of a DUQL query
// type Query struct {
// 	Settings *Settings      `yaml:"settings,omitempty" json:"settings,omitempty" mapstructure:"settings,omitempty"`
// 	Declare  map[string]any `yaml:"declare,omitempty" json:"declare,omitempty" mapstructure:"declare,omitempty"`
// 	Dataset  Dataset        `yaml:"dataset" json:"dataset" mapstructure:"dataset"`
// 	Steps    []Step         `yaml:"steps,omitempty" json:"steps,omitempty" mapstructure:"steps,omitempty"`
// 	Into     string         `yaml:"into,omitempty" json:"into,omitempty" mapstructure:"into,omitempty"`
// }

// // Settings represents metadata and configuration options for DUQL queries
// type Settings struct {
// 	Version string        `yaml:"version,omitempty" json:"version,omitempty" mapstructure:"version,omitempty"`
// 	Target  TargetDialect `yaml:"target,omitempty" json:"target,omitempty" mapstructure:"target,omitempty"`
// }

// // TargetDialect represents the supported SQL dialects
// type TargetDialect string

// const (
// 	ClickHouse TargetDialect = "sql.clickhouse"
// 	DuckDB     TargetDialect = "sql.duckdb"
// 	Generic    TargetDialect = "sql.generic"
// 	GlareDB    TargetDialect = "sql.glaredb"
// 	MySQL      TargetDialect = "sql.mysql"
// 	Postgres   TargetDialect = "sql.postgres"
// 	SQLite     TargetDialect = "sql.sqlite"
// )

// // Dataset represents the source of data for a DUQL query
// type Dataset struct {
// 	Name   string     `yaml:"name,omitempty" json:"name,omitempty" mapstructure:"name,omitempty"`
// 	Format DataFormat `yaml:"format,omitempty" json:"format,omitempty" mapstructure:"format,omitempty"`
// }

// // DataFormat represents the supported data formats
// type DataFormat string

// const (
// 	Table   DataFormat = "table"
// 	CSV     DataFormat = "csv"
// 	JSON    DataFormat = "json"
// 	Parquet DataFormat = "parquet"
// )

// Step represents a single operation in the DUQL query pipeline
// type Step struct {
// 	Filter    *Expression            `yaml:"filter,omitempty" json:"filter,omitempty" mapstructure:"filter,omitempty"`
// 	Generate  map[string]interface{} `yaml:"generate,omitempty" json:"generate,omitempty" mapstructure:"generate,omitempty"`
// 	Group     *Group                 `yaml:"group,omitempty" json:"group,omitempty" mapstructure:"group,omitempty"`
// 	Join      *Join                  `yaml:"join,omitempty" json:"join,omitempty" mapstructure:"join,omitempty"`
// 	Select    *Select                `yaml:"select,omitempty" json:"select,omitempty" mapstructure:"select,omitempty"`
// 	SelectNot []string               `yaml:"select!,omitempty" json:"select!,omitempty" mapstructure:"select!,omitempty"`
// 	Sort      *Sort                  `yaml:"sort,omitempty" json:"sort,omitempty" mapstructure:"sort,omitempty"`
// 	Take      *Take                  `yaml:"take,omitempty" json:"take,omitempty" mapstructure:"take,omitempty"`
// 	Window    *Window                `yaml:"window,omitempty" json:"window,omitempty" mapstructure:"window,omitempty"`
// 	Loop      []Step                 `yaml:"loop,omitempty" json:"loop,omitempty" mapstructure:"loop,omitempty"`
// }

// // Expression represents a DUQL expression, which can be a string or a complex object

// // Group represents a grouping operation in DUQL

// // Join represents a join operation in DUQL

// // JoinType represents the supported join types

// // Select represents the select operation in DUQL

// // Sort represents the sort operation in DUQL

// // Take represents the take operation in DUQL

// // Window represents a window function in DUQL

// // UnmarshalYAML implements custom YAML unmarshaling for Dataset
// func (d *Dataset) UnmarshalYAML(value *yaml.Node) error {
// 	var s string
// 	if err := value.Decode(&s); err == nil {
// 		d.Name = s
// 		return nil
// 	}

// 	type datasetAlias Dataset
// 	return value.Decode((*datasetAlias)(d))
// }

// // UnmarshalJSON implements custom JSON unmarshaling for Dataset
// func (d *Dataset) UnmarshalJSON(data []byte) error {
// 	var s string
// 	if err := json.Unmarshal(data, &s); err == nil {
// 		d.Name = s
// 		return nil
// 	}

// 	type datasetAlias Dataset
// 	return json.Unmarshal(data, (*datasetAlias)(d))
// }

// // CreateDUQLQuery creates a new DUQL query structure
// func CreateDUQLQuery() *Query {
// 	return &Query{
// 		Settings: &Settings{},
// 		Declare:  make(map[string]any),
// 		Steps:    []Step{},
// 	}
// }
