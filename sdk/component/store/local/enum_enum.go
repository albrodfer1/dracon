// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package localstore

import (
	"errors"
	"fmt"
)

const (
	// LocalStoreColumnNameFindings is a localStoreColumnName of type findings.
	LocalStoreColumnNameFindings localStoreColumnName = "findings"
	// LocalStoreColumnNameInstanceId is a localStoreColumnName of type instance_id.
	LocalStoreColumnNameInstanceId localStoreColumnName = "instance_id"
	// LocalStoreColumnNameUpdatedAt is a localStoreColumnName of type updated_at.
	LocalStoreColumnNameUpdatedAt localStoreColumnName = "updated_at"
)

var ErrInvalidlocalStoreColumnName = errors.New("not a valid localStoreColumnName")

// String implements the Stringer interface.
func (x localStoreColumnName) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x localStoreColumnName) IsValid() bool {
	_, err := ParselocalStoreColumnName(string(x))
	return err == nil
}

var _localStoreColumnNameValue = map[string]localStoreColumnName{
	"findings":    LocalStoreColumnNameFindings,
	"instance_id": LocalStoreColumnNameInstanceId,
	"updated_at":  LocalStoreColumnNameUpdatedAt,
}

// ParselocalStoreColumnName attempts to convert a string to a localStoreColumnName.
func ParselocalStoreColumnName(name string) (localStoreColumnName, error) {
	if x, ok := _localStoreColumnNameValue[name]; ok {
		return x, nil
	}
	return localStoreColumnName(""), fmt.Errorf("%s is %w", name, ErrInvalidlocalStoreColumnName)
}
