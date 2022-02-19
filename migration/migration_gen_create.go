package migration

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const createTableRaw = `CREATE {{.temporary}} TABLE {{.tableName}}(
{{.fields}}
) {{.engine}} {{.charset}};`

var (
	// 需要带长度的字段类型
	typeLenMap = map[string]string{
		"char":"char",
		"varchar":"varchar",
		"string":"varchar",
	}
	// 需要带精度的类型
	typePrecisionMap = map[string]string{
		"decimal":"decimal",
		"float":"float",
		"double":"double",
	}
)

func (m *MigrationTable) generateCreateMigrationSQL() {
	raw := createTableRaw
	raw = strings.Replace(raw,"{{.tableName}}",m.table.tableName,1)
	//
	engine := ""
	if m.table.engine != "" {
		engine = "ENGINE="+m.table.engine
	}
	raw = strings.Replace(raw,"{{.engine}}",engine,1)
	//
	charset := ""
	if m.table.charset != "" {
		charset = "CHARSET="+m.table.charset
	}
	raw = strings.Replace(raw,"{{.charset}}",charset,1)
	//
	fields := []string{}
	primarys := []string{}
	for _,v := range m.table.result {
		if v.fieldType == "timestamps" {
			v.field = "created_at"
			v.fieldType = "timestamp"
			v.canNull = true
			fields = append(fields,m.generateCreateMigration(v))
			v.field = "updated_at"
		}
		fields = append(fields,m.generateCreateMigration(v))
		if v.primaryKey {
			primarys = append(primarys,v.field)
		}
	}
	// primary
	if len(primarys) != 0 {
		for k,v := range primarys {
			primarys[k] = "`" + v + "`"
		}
		fields = append(fields,fmt.Sprintf("PRIMARY KEY(%v)",strings.Join(primarys,",")))
	}
	// unique
	if len(m.table.uniqueName) != 0 && len(m.table.uniqueFields) != 0 {
		for i := 0; i < len(m.table.uniqueFields); i++ {
			for k,v := range m.table.uniqueFields[i] {
				m.table.uniqueFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("UNIQUE KEY `%v` (%v)  ",m.table.uniqueName[i],strings.Join(m.table.uniqueFields[i],",")))
		}
	}
	// index
	if len(m.table.indexName) != 0 && len(m.table.indexFields) != 0 {
		for i:=0;i<len(m.table.indexFields);i ++ {
			for k,v := range m.table.indexFields[i] {
				m.table.indexFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("KEY `%v` (%v)  ",m.table.indexName[i],strings.Join(m.table.indexFields[i],",")))
		}
	}
	// spatial index
	if len(m.table.spatialName) != 0 && len(m.table.spatialFields) != 0 {
		for i:=0;i<len(m.table.spatialFields);i ++ {
			for k,v := range m.table.spatialFields[i] {
				m.table.spatialFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("SPATIAL INDEX `%v` (%v)",m.table.spatialName[i],strings.Join(m.table.spatialFields[i],",")))
		}
	}
	// foreign
	if m.table.foreign != nil && len(m.table.foreign.result) != 0 {
		for _,v := range m.table.foreign.result {
			fields = append(fields,fmt.Sprintf("foreign key (`%v`) references %v(`%v`)",v.key,v.referencesTable,v.referencesField))
			if v.onUpdate {
				fields[len(fields)-1] += " on update cascade "
			}
			if v.onDelete {
				fields[len(fields)-1] += " on delete cascade "
			}
		}
	}
	temp := ""
	if m.table.temporary {
		temp = "temporary"
	}
	raw = strings.Replace(raw,"{{.temporary}}",temp,1)
	//
	fieldRaw := strings.Join(fields,",\n")
	raw = strings.Replace(raw,"{{.fields}}",fieldRaw,1)
	fmt.Println(raw)
}

// 生成表中的字段
func (m *MigrationTable) generateCreateMigration(t *MigrationAttribute) string {
	field := ""
	// field name
	field += "`"+t.field+"` "
	// type
	fieldType := t.fieldType
	t.fieldType = strings.ToLower(t.fieldType)
	if k,ok := typeLenMap[t.fieldType];ok {
		if t.length == 0 {
			t.length = 255
		}
		fieldType = k + fmt.Sprintf("(%d)",t.length)
	} else if k,ok := typePrecisionMap[t.fieldType];ok {
		if t.precision < 1 {
			t.precision = 0
		}
		if t.length < 1 {
			t.length = 10
		}
		fieldType = k + fmt.Sprintf("(%d,%d)",t.length,t.precision)
	}
	field += fieldType
	if t.unsigned {
		field += " UNSIGNED "
	}
	// not null
	if !t.canNull {
		field += " NOT NULL "
	}
	if t.canNull {
		field += " NULL "
	}
	if t.defaultValue != nil {
		field += " " + defaultValue(t.defaultValue)
	}
	if t.comment != "" {
		field += " comment '" + t.comment + "'"
	}
	if t.autoIncrement {
		field += " AUTO_INCREMENT "
	}
	if t.unique {
		field += " UNIQUE"
	}
	return field
}

func defaultValue(d interface{}) string {
	switch reflect.TypeOf(d).Kind() {
	case reflect.String:
		if d.(string) == "CURRENT_TIMESTAMP" {
			return "DEFAULT CURRENT_TIMESTAMP"
		}
		return "DEFAULT '" + d.(string) + "'"
	case reflect.Int:
		return "DEFAULT " + strconv.FormatInt(int64(d.(int)),10)
	case reflect.Bool:
		if d.(bool) {
			return "DEFAULT 0"
		}
		return "DEFAULT 1"
	default:
		return ""
	}
}