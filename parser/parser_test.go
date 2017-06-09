package parser

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParsePackageReturnsErrorOnEmptyDirectoryString(t *testing.T) {
	_, err := ParsePackage("")
	assert.Error(t, err)
}

func TestParsePackageReturnsErrorDirectoryUp(t *testing.T) {
	_, err := ParsePackage("..")
	assert.Error(t, err)
}

func TestParsePackageReturnsPackageWhenGivenAValidPackage(t *testing.T) {
	dir, err := filepath.Abs("../test")
	assert.NoError(t, err)
	pkg, err := ParsePackage(dir)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
}

func TestPackage_ValuesOfTypeReturnsErrorWithFakeConstantType(t *testing.T) {
	dir, err := filepath.Abs("../test")
	assert.NoError(t, err)
	pkg, err := ParsePackage(dir)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)

	values, err := pkg.ValuesOfType("cars")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no values defined for type")
	assert.Nil(t, values)
}

func TestPackage_ValuesOfTypeErrorsWithNonIntegerConstant(t *testing.T) {
	dir, err := filepath.Abs("../test")
	assert.NoError(t, err)
	pkg, err := ParsePackage(dir)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)

	values, err := pkg.ValuesOfType("bar")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't handle non-integer")
	assert.Nil(t, values)
}

// Validate that we get all values from the constants of a given type
// This test also validates that _ is ignored when parsing the code
func TestPackage_ValuesOfTypeReturnsProperValues(t *testing.T) {
	dir, err := filepath.Abs("../test")
	assert.NoError(t, err)
	pkg, err := ParsePackage(dir)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)

	values, err := pkg.ValuesOfType("ShirtSize")
	assert.NoError(t, err)
	assert.Len(t, values, 6)
	// validate with one of the values
	assert.Contains(t, values, "M")
}
