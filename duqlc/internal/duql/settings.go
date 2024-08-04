package duql

type Settings struct {
	Version string        `yaml:"version,omitempty" json:"version,omitempty" mapstructure:"version,omitempty"`
	Target  TargetDialect `yaml:"target,omitempty" json:"target,omitempty" mapstructure:"target,omitempty"`
}

type TargetDialect string

const (
	ClickHouse TargetDialect = "sql.clickhouse"
	DuckDB     TargetDialect = "sql.duckdb"
	Generic    TargetDialect = "sql.generic"
	GlareDB    TargetDialect = "sql.glaredb"
	MySQL      TargetDialect = "sql.mysql"
	Postgres   TargetDialect = "sql.postgres"
	SQLite     TargetDialect = "sql.sqlite"
)
