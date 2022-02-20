package migration

type migrationer interface {
	Up()

	Down()
}

type MigrationAttributes struct {
	tableName     string
	engine        string
	charset       string
	collation     string
	temporary     bool
	result        []*MigrationAttribute
	uniqueName    []string
	uniqueFields  [][]string
	indexName     []string
	indexFields   [][]string
	spatialName   []string
	spatialFields [][]string
	foreign       *ForeignAttributes
}

type ForeignAttributes struct {
	result []*foreignAttribute
}

type foreignAttribute struct {
	key             string
	referencesTable string
	referencesField string
	onUpdate        bool
	onDelete        bool
}

type MigrationAttribute struct {
	field           string
	fieldType       string
	length          int
	precision       int
	defaultValue    interface{}
	canNull         bool
	comment         string
	maintenanceTime bool
	rename          string
	dropField       []string
	unique          bool
	references      string
	referencesOn    string
	isChange        bool
	primaryKey      bool
	autoIncrement   bool
	unsigned        bool
	index           bool
}

type MigrationForegion struct {
	table *MigrationAttributes
}

type MigrationTable struct {
	table *MigrationAttributes
}
