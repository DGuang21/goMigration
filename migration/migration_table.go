package migration

type MigrationTable struct {
	*MigrationAttributes
	result []*MigrationAttributes
}

/**
 * @Description: 递增 ID（主键），相当于「UNSIGNED BIG INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) BigIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "bigIncrements",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) BigInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "bigInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Binary(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "binary",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Boolean(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "boolean",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Char(field string,length ...int) *MigrationAttributes {
	m.appendLastResult()
	l := 0
	if len(length) != 0 {
		l = length[0]
	}
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "char",
		length: l,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) date(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "date",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) dateTime(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "dateTime",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) dateTimeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "dateTimeTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) decimal(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "decimal",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) double(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "double",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) float(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "double",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) geometry(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "geometry",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) geometryCollection(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "geometryCollection",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) increments(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "increments",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) integer(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "integer",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) ipAddress(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "ipAddress",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) json(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "json",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) jsonb(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "jsonb",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) lineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "lineString",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) longText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "longText",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) macAddress(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "macAddress",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) mediumIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "mediumIncrements",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) mediumInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "mediumInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) mediumText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "mediumText",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) morphs(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "morphs",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) multiLineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "multiLineString",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) multiPoint(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "multiPoint",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) multiPolygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "multiPolygon",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) nullableMorphs(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "nullableMorphs",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) nullableTimestamps(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "nullableTimestamps",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) point(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "point",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) polygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "polygon",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) smallIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "smallIncrements",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) smallInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "smallInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) softDeletes() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "delete_at",
		fieldType:       "softDeletes",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) softDeletesTz() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "delete_at",
		fieldType:       "softDeletesTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) text(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "text",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) time(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "time",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) timeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timeTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) timestamp(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) timestampTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestampTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Timestamps() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "",
		fieldType:       "timestampTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) timestampsTz() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "",
		fieldType:       "timestampsTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) tinyIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyIncrements",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) tinyInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedBigInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedBigInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedDecimal(field string,length,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedBigInteger",
		length:     length,
		precision: precision,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedMediumInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedMediumInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) unsignedSmallInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedSmallInteger",
	}
	return m.MigrationAttributes
}


func (m *MigrationTable) unsignedTinyInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "unsignedTinyInteger",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) uuid(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "uuid",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) year(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "year",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) comment(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "comment",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) String(field string,length ...int) *MigrationAttributes {
	m.appendLastResult()
	l := 0
	if len(length) != 0 {
		l = length[0]
	}
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "string",
		length:          l,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Done() {
	m.appendLastResult()
}

func (m *MigrationTable) appendLastResult() {
	if m.MigrationAttributes != nil {
		m.result = append(m.result,m.MigrationAttributes)
	}
}