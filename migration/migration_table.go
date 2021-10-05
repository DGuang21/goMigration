package migration

type MigrationTable struct {
	*MigrationAttributes
	result []*MigrationAttributes
	tableName       string
	engine          string
	charset   		string
	collation		string
	temporary		bool
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
		fieldType:       "UNSIGNED BIG INTEGER",
		primaryKey:      true,
		autoIncrement:   true,
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于 BIGINT
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) BigInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "BIGINT",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于 BLOB
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Binary(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "BLOB",
	}
	return m.MigrationAttributes
}

/**
 * @Description:
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Boolean(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "tinyint",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于带有长度的 CHAR
 * @receiver m
 * @param field
 * @param length
 * @auth: daguang
 * @return *MigrationAttributes
 */
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

func (m *MigrationTable) Date(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DATE",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) DateTime(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DATETIME",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) DateTimeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Decimal(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DECIMAL",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Double(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "DOUBLE",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Float(field string,length int,precision int) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "FLOAT",
		precision: precision,
		length: length,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Geometry(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "GEOMETRY",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) GeometryCollection(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "GEOMETRYCOLLECTION",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增的 ID (主键)，相当于「UNSIGNED INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) Increments(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "UNSIGNED INTEGER",
		primaryKey: true,
		autoIncrement: true,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Integer(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "INTEGER",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Json(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "JSON",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Jsonb(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "JSONB",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) LineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "LINESTRING",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) LongText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "LONGTEXT",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增 ID (主键) ，相当于「UNSIGNED MEDIUM INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) MediumIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "UNSIGNED MEDIUM INTEGER",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MediumInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MEDIUMINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MediumText(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MEDIUMTEXT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiLineString(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTILINESTRING",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPoint(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTIPOINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) MultiPolygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "MULTIPOLYGON",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) NullableTimestamps(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamps",
		canNull: true,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Point(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "POINT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Polygon(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "POLYGON",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 递增 ID (主键) ，相当于「UNSIGNED SMALL INTEGER」
 * @receiver m
 * @param field
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) SmallIncrements(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "UNSIGNED SMALL INTEGER",
		primaryKey: true,
		autoIncrement: true,
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) SmallInteger(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "SMALLINT",
	}
	return m.MigrationAttributes
}

/**
 * @Description: 相当于为软删除添加一个可空的 deleted_at 字段
 * @receiver m
 * @auth: daguang
 * @return *MigrationAttributes
 */
func (m *MigrationTable) SoftDeletes() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "delete_at",
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Text(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "TEXT",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Time(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "TIME",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimeTz(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timeTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) Timestamp(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "timestamp",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimestampTz(field string) *MigrationAttributes {
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
		fieldType:       "timestamps",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TimestampsTz() *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           "",
		fieldType:       "timestampsTz",
	}
	return m.MigrationAttributes
}

func (m *MigrationTable) TinyIncrements(field string) *MigrationAttributes {
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

func (m *MigrationTable) UUID(field string) *MigrationAttributes {
	m.appendLastResult()
	m.MigrationAttributes = &MigrationAttributes{
		field:           field,
		fieldType:       "uuid",
	}
	return m.MigrationAttributes
}

// Comment 表注释
func (m *MigrationTable) Comment(field string) *MigrationAttributes {
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

func (m *MigrationTable) Engine(engine string) {
	m.engine = engine
}

func (m *MigrationTable) Charset(charset string) {
	m.charset = charset
}

func (m *MigrationTable) Collation(collation string) {
	m.collation = collation
}

func (m *MigrationTable) Temporary() {
	m.temporary = true
}

func (m *MigrationTable) Done() {
	m.appendLastResult()
}

func (m *MigrationTable) appendLastResult() {
	if m.MigrationAttributes != nil {
		m.result = append(m.result,m.MigrationAttributes)
	}
}