//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var DynamicChallenge = newDynamicChallengeTable("public", "dynamic_challenge", "")

type dynamicChallengeTable struct {
	postgres.Table

	//Columns
	ID      postgres.ColumnString
	Initial postgres.ColumnInteger
	Minimum postgres.ColumnInteger
	Decay   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type DynamicChallengeTable struct {
	dynamicChallengeTable

	EXCLUDED dynamicChallengeTable
}

// AS creates new DynamicChallengeTable with assigned alias
func (a DynamicChallengeTable) AS(alias string) *DynamicChallengeTable {
	return newDynamicChallengeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DynamicChallengeTable with assigned schema name
func (a DynamicChallengeTable) FromSchema(schemaName string) *DynamicChallengeTable {
	return newDynamicChallengeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DynamicChallengeTable with assigned table prefix
func (a DynamicChallengeTable) WithPrefix(prefix string) *DynamicChallengeTable {
	return newDynamicChallengeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DynamicChallengeTable with assigned table suffix
func (a DynamicChallengeTable) WithSuffix(suffix string) *DynamicChallengeTable {
	return newDynamicChallengeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDynamicChallengeTable(schemaName, tableName, alias string) *DynamicChallengeTable {
	return &DynamicChallengeTable{
		dynamicChallengeTable: newDynamicChallengeTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newDynamicChallengeTableImpl("", "excluded", ""),
	}
}

func newDynamicChallengeTableImpl(schemaName, tableName, alias string) dynamicChallengeTable {
	var (
		IDColumn       = postgres.StringColumn("id")
		InitialColumn  = postgres.IntegerColumn("initial")
		MinimumColumn  = postgres.IntegerColumn("minimum")
		DecayColumn    = postgres.IntegerColumn("decay")
		allColumns     = postgres.ColumnList{IDColumn, InitialColumn, MinimumColumn, DecayColumn}
		mutableColumns = postgres.ColumnList{InitialColumn, MinimumColumn, DecayColumn}
	)

	return dynamicChallengeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		Initial: InitialColumn,
		Minimum: MinimumColumn,
		Decay:   DecayColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
