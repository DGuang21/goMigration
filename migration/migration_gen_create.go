package migration

import (
	"fmt"
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
	typePrecisionMap = map[string]string{}
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
	for _,v := range m.result {
		if v.fieldType == "timestamps" {
			v.field = "created_at"
			v.fieldType = "timestamp"
			v.canNull = true
			fields = append(fields,m.generateCreateMigration(v))
			v.field = "updated_at"
		}
		fields = append(fields,m.generateCreateMigration(v))
	}
	//
	temp := ""
	if m.temporary {
		temp = "temporary"
	}
	raw = strings.Replace(raw,"{{.temporary}}",temp,1)
	//
	raw = strings.Replace(raw,"{{.fields}}",strings.Join(fields,",\n"),1)
	fmt.Println(raw)
}

// 生成表中的字段
func (m *MigrationTable) generateCreateMigration(t *MigrationAttribute) string {
	field := ""
	// field name
	field += "`"+t.field+"` "
	// type
	fieldType := t.fieldType
	if k,ok := typeLenMap[t.fieldType];ok {
		if t.length == 0 {
			t.length = 255
		}
		fieldType = k + fmt.Sprintf("(%d)",t.length)
	}
	field += fieldType
	// not null
	if !t.canNull {
		field += " NOT NULL "
	}
	if t.canNull {
		field += " NULL "
	}
	return field
}