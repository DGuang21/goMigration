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
	raw = strings.Replace(raw,"{{.tableName}}",m.tableName,1)
	//
	engine := ""
	if m.engine != "" {
		engine = "ENGINE="+m.engine
	}
	raw = strings.Replace(raw,"{{.engine}}",engine,1)
	//
	charset := ""
	if m.charset != "" {
		charset = "CHARSET="+m.charset
	}
	raw = strings.Replace(raw,"{{.charset}}",charset,1)
	//
	fields := []string{}
	primarys := []string{}
	for _,v := range m.result {
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
	if len(m.uniqueName) != 0 && len(m.uniqueFields) != 0 {
		for i := 0; i < len(m.uniqueFields); i++ {
			for k,v := range m.uniqueFields[i] {
				m.uniqueFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("UNIQUE KEY `%v` (%v)  ",m.uniqueName[i],strings.Join(m.uniqueFields[i],",")))
		}
	}
	// index
	if len(m.indexName) != 0 && len(m.indexFields) != 0 {
		for i:=0;i<len(m.indexFields);i ++ {
			for k,v := range m.indexFields[i] {
				m.indexFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("KEY `%v` (%v)  ",m.indexName[i],strings.Join(m.indexFields[i],",")))
		}
	}
	// spatial index
	if len(m.spatialName) != 0 && len(m.spatialFields) != 0 {
		for i:=0;i<len(m.spatialFields);i ++ {
			for k,v := range m.spatialFields[i] {
				m.spatialFields[i][k] = "`" + v + "`"
			}
			fields = append(fields,fmt.Sprintf("SPATIAL INDEX `%v` (%v)",m.spatialName[i],strings.Join(m.spatialFields[i],",")))
		}
	}
	//
	temp := ""
	if m.temporary {
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