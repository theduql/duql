package validator

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	duql "github.com/theduql/duql/internal/duql"
	"github.com/theduql/duql/internal/logger"
)

func init() {
	// schemaDir := "../schema"
	// rootSchemaPath := filepath.Join(schemaDir, "query.s.duql.json")

	// // Create a reference loader that can resolve references
	// schemaLoader = gojsonschema.NewReferenceLoader("file://" + rootSchemaPath)

	// // Optionally, validate that the schema itself is valid
	// _, err := gojsonschema.NewSchema(schemaLoader)
	// if err != nil {
	// 	log.Fatalf("Failed to load schema: %v", err)
	// }
}

func Validate(path string) error {
	log := logger.GetLogger()

	info, err := os.Stat(path)
	if err != nil {
		log.Error(fmt.Sprintf("❌ Unable to Access Provided Path: %s", path))
		return err
	}

	if info.IsDir() {
		return validateDir(path)
	}

	return validateFile(path)
}

func validateFile(file string) error {
	log := logger.GetLogger()

	log.Info(fmt.Sprintf("ℹ️  Validating File: %s", file))

	// Read the file
	data, err := os.ReadFile(file)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to Read File: %s", err))
		return err
	}
	log.Info(fmt.Sprintf("✅ Opened File: %s", file))
	log.Debug("File contents", zap.String("content", string(data)))

	// Attempt to unmarshal YAML into DUQL Query structure
	var query duql.Query
	err = yaml.Unmarshal(data, &query)
	if err != nil {
		log.Error(fmt.Sprintf("Invalid DUQL Query: %s", err))
		return err
	}

	// Validate the query
	if err := query.Validate(); err != nil {
		log.Error(fmt.Sprintf("Invalid DUQL Query: %s", err))
		return err
	}

	log.Info("✅ Valid YAML and conforms to DUQL schema")

	return nil
}

func validateDir(dir string) error {
	log := logger.GetLogger()

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error("Error walking directory", zap.String("path", path), zap.Error(err))
			return fmt.Errorf("error walking directory: %s", err.Error())
		}
		if !info.IsDir() && (filepath.Ext(path) == ".duql" || filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
			if err := validateFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}
