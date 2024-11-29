package repository

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

// BaseRepository provides common database operations and utilities
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

// DB returns the underlying database connection
func (r *BaseRepository) DB() *gorm.DB {
	return r.db
}

// Pagination represents pagination parameters
type Pagination struct {
	Limit  int
	Offset int
}

// GetPagination returns limit and offset based on the provided filter
func (r *BaseRepository) GetPagination(filter map[string]interface{}) Pagination {
	limit := 25 // default limit
	if val, ok := filter["limit"]; ok {
		if l, ok := val.(int); ok {
			limit = l
		}
	}

	offset := 0
	if val, ok := filter["offset"]; ok {
		if o, ok := val.(int); ok {
			if o > 1 {
				offset = ((o - 1) * limit) + 1
			} else if o == 1 {
				offset = 0
			}
		}
	}

	return Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

// Save saves the model to the database
func (r *BaseRepository) Save(model interface{}) error {
	return r.db.Save(model).Error
}

// DeleteByField deletes records by a specific field value
func (r *BaseRepository) DeleteByField(model interface{}, deleteIDs []interface{}, field string) ([]uint, error) {
	if len(deleteIDs) == 0 {
		return nil, nil
	}

	var processedIDs []uint
	result := r.db.Model(model).Where(fmt.Sprintf("%s IN ?", field), deleteIDs).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	// Extract IDs from the deleted records
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if id := item.FieldByName("ID"); id.IsValid() {
				processedIDs = append(processedIDs, uint(id.Uint()))
			}
		}
	}

	// Perform the deletion
	if err := r.db.Delete(model, fmt.Sprintf("%s IN ?", field), deleteIDs).Error; err != nil {
		return nil, err
	}

	return processedIDs, nil
}

// DeleteIDNotIn deletes records where ID is not in the provided list
func (r *BaseRepository) DeleteIDNotIn(model interface{}, ids []uint, conditions map[string]interface{}) error {
	query := r.db.Model(model)
	if len(ids) > 0 {
		query = query.Where("id NOT IN ?", ids)
	}
	
	// Apply additional conditions
	for key, value := range conditions {
		query = query.Where(key, value)
	}

	return query.Delete(model).Error
}

// MapDataToModel maps data from map to model fields
func (r *BaseRepository) MapDataToModel(model interface{}, data map[string]interface{}, userID uint, fieldList []string) error {
	now := time.Now()
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// If fieldList is provided, only map specified fields
	if len(fieldList) > 0 {
		for _, field := range fieldList {
			if value, ok := data[field]; ok {
				if field != "id" && !isComplexType(value) {
					if f := val.FieldByName(strings.Title(field)); f.IsValid() && f.CanSet() {
						if err := setField(f, value); err != nil {
							return err
						}
					}
				}
			}
		}
	} else {
		// Map all fields
		for key, value := range data {
			if key != "id" && !isComplexType(value) {
				if f := val.FieldByName(strings.Title(key)); f.IsValid() && f.CanSet() {
					if err := setField(f, value); err != nil {
						return err
					}
				}
			}
		}
	}

	// Set audit fields
	if f := val.FieldByName("UpdatedBy"); f.IsValid() && f.CanSet() {
		f.SetUint(uint64(userID))
	}
	if f := val.FieldByName("UpdatedAt"); f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(now))
	}

	// Set created fields if they're empty
	if f := val.FieldByName("CreatedBy"); f.IsValid() && f.CanSet() && f.IsZero() {
		f.SetUint(uint64(userID))
	}
	if f := val.FieldByName("CreatedAt"); f.IsValid() && f.CanSet() && f.IsZero() {
		f.Set(reflect.ValueOf(now))
	}

	return nil
}

// BuildQuery builds a query with the provided conditions
func (r *BaseRepository) BuildQuery(query *gorm.DB, conditions map[string]interface{}) *gorm.DB {
	for key, value := range conditions {
		if v, ok := value.(map[string]interface{}); ok {
			if op, ok := v["Op"].(string); ok {
				switch op {
				case "eq":
					query = query.Where(fmt.Sprintf("%s = ?", key), v["value"])
				case "ne":
					query = query.Where(fmt.Sprintf("%s != ?", key), v["value"])
				case "in":
					query = query.Where(fmt.Sprintf("%s IN ?", key), v["value"])
				case "notIn":
					query = query.Where(fmt.Sprintf("%s NOT IN ?", key), v["value"])
				case "gt":
					query = query.Where(fmt.Sprintf("%s > ?", key), v["value"])
				case "gte":
					query = query.Where(fmt.Sprintf("%s >= ?", key), v["value"])
				case "lt":
					query = query.Where(fmt.Sprintf("%s < ?", key), v["value"])
				case "lte":
					query = query.Where(fmt.Sprintf("%s <= ?", key), v["value"])
				case "between":
					if vals, ok := v["value"].([]interface{}); ok && len(vals) == 2 {
						query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", key), vals[0], vals[1])
					}
				}
			}
		} else {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	return query
}

// ToJSON converts a model or slice of models to JSON
func (r *BaseRepository) ToJSON(model interface{}) ([]byte, error) {
	return json.Marshal(model)
}

// NotFound checks if a record exists and returns an error if it doesn't
func (r *BaseRepository) NotFound(model interface{}, msg string) error {
	if reflect.ValueOf(model).IsNil() {
		if msg == "" {
			msg = "Record not found"
		}
		return fmt.Errorf(msg)
	}
	return nil
}

// Helper functions

func isComplexType(v interface{}) bool {
	switch v.(type) {
	case map[string]interface{}, []interface{}:
		return true
	default:
		return false
	}
}

func setField(field reflect.Value, value interface{}) error {
	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("invalid type for field")
	}
	field.Set(val)
	return nil
}
