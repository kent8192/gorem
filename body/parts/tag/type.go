package tag

import "fmt"

// DataTypeInterface
type DataTypeInterface interface {
	Refer() string
}

// DataType構造体
type DataType struct {
	dataType                  DataTypeInterface
	isLegitimatelyInitialized bool
}

// InitTagメソッド
func (d *DataType) InitTag(value DataTypeInterface) *DataType {
	d.isLegitimatelyInitialized = true
	d.dataType = value
	if err := d.ValidateTag(); err != nil {
		panic(err)
	}
	return d
}

// RenderTagメソッド
func (d *DataType) RenderTag() string {
	if !d.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("type:%s", d.dataType.Refer())
}

// IsValidメソッド
func (d *DataType) IsValid() (bool, string) {
	if !d.isLegitimatelyInitialized {
		return false, "data type is not initialized legitimately"
	}
	if d.dataType == nil {
		return false, "data type interface is nil"
	}
	if d.dataType.Refer() == "" {
		return false, "data type's Refer method returned an empty string"
	}
	return true, ""
}

// ValidateTagメソッド
func (d *DataType) ValidateTag() error {
	// エラー内容をIsValidから受け取る
	if valid, errMessage := d.IsValid(); !valid {
		return fmt.Errorf("invalid data type: %s", errMessage)
	}
	return nil
}

func (r *DataType) GetIsLegitimatelyInitialized() bool {
	return r.isLegitimatelyInitialized
}

// MySQLDataTypeは、GORMのtypeタグで使用するMySQLのデータ型を表す列挙型です。
type MySQLDataType int

const (
	// 数値型
	MySQLTinyInt MySQLDataType = iota
	MySQLSmallInt
	MySQLMediumInt
	MySQLInt
	MySQLBigInt
	MySQLDecimal
	MySQLNumeric
	MySQLFloat
	MySQLDouble
	// 文字列型
	MySQLChar
	MySQLVarChar
	MySQLTinyText
	MySQLText
	MySQLMediumText
	MySQLLongText
	// バイナリ型
	MySQLBinary
	MySQLVarBinary
	MySQLTinyBlob
	MySQLBlob
	MySQLMediumBlob
	MySQLLongBlob
	// 日時型
	MySQLDate
	MySQLTime
	MySQLDateTime
	MySQLTimestamp
	MySQLYear
	// その他の型
	MySQLEnum
	MySQLSet
	MySQLJSON
)

// Stringは、MySQLDataTypeの値を文字列として返します。
func (dt MySQLDataType) Refer() string {
	return [...]string{
		"tinyint", "smallint", "mediumint", "int", "bigint",
		"decimal", "numeric", "float", "double",
		"char", "varchar", "tinytext", "text", "mediumtext", "longtext",
		"binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob",
		"date", "time", "datetime", "timestamp", "year",
		"enum", "set", "json",
	}[dt]
}

// PostgreSQLDataTypeは、GORMのtypeタグで指定するPostgreSQLのデータ型を表す列挙型です。
type PostgreSQLDataType int

const (
	PostgreSQLBigInt PostgreSQLDataType = iota
	PostgreSQLBigSerial
	PostgreSQLBit
	PostgreSQLBitVarying
	PostgreSQLBoolean
	PostgreSQLBox
	PostgreSQLBytea
	PostgreSQLCharacter
	PostgreSQLCharacterVarying
	PostgreSQLCIDR
	PostgreSQLCircle
	PostgreSQLDate
	PostgreSQLDoublePrecision
	PostgreSQLInet
	PostgreSQLInteger
	PostgreSQLInterval
	PostgreSQLJSON
	PostgreSQLJSONB
	PostgreSQLLine
	PostgreSQLLseg
	PostgreSQLMacaddr
	PostgreSQLMacaddr8
	PostgreSQLMoney
	PostgreSQLNumeric
	PostgreSQLPath
	PostgreSQLPGLSN
	PostgreSQLPoint
	PostgreSQLPolygon
	PostgreSQLReal
	PostgreSQLSmallInt
	PostgreSQLSmallSerial
	PostgreSQLSerial
	PostgreSQLText
	PostgreSQLTime
	PostgreSQLTimeWithTimeZone
	PostgreSQLTimestamp
	PostgreSQLTimestampWithTimeZone
	PostgreSQLTSQuery
	PostgreSQLTSVector
	PostgreSQLTxidSnapshot
	PostgreSQLUUID
	PostgreSQLXML
)

// Stringメソッドは、PostgreSQLDataTypeの値を文字列に変換します。
func (dt PostgreSQLDataType) Refer() string {
	return [...]string{
		"bigint",
		"bigserial",
		"bit",
		"bit varying",
		"boolean",
		"box",
		"bytea",
		"character",
		"character varying",
		"cidr",
		"circle",
		"date",
		"double precision",
		"inet",
		"integer",
		"interval",
		"json",
		"jsonb",
		"line",
		"lseg",
		"macaddr",
		"macaddr8",
		"money",
		"numeric",
		"path",
		"pg_lsn",
		"point",
		"polygon",
		"real",
		"smallint",
		"smallserial",
		"serial",
		"text",
		"time",
		"time with time zone",
		"timestamp",
		"timestamp with time zone",
		"tsquery",
		"tsvector",
		"txid_snapshot",
		"uuid",
		"xml",
	}[dt]
}

// SQLiteDataTypeは、GORMのtypeタグで指定するSQLiteのデータ型を表す列挙型です。
type SQLiteDataType int

const (
	SQLiteInteger SQLiteDataType = iota
	SQLiteReal
	SQLiteText
	SQLiteBlob
	SQLiteNumeric
)

// Stringメソッドは、SQLiteDataTypeの値を文字列に変換します。
func (dt SQLiteDataType) Refer() string {
	return [...]string{
		"INTEGER",
		"REAL",
		"TEXT",
		"BLOB",
		"NUMERIC",
	}[dt]
}

// OracleDataTypeは、GORMのtypeタグで指定するOracleのデータ型を表す列挙型です。
type OracleDataType int

const (
	OracleChar OracleDataType = iota
	OracleVarchar2
	OracleNchar
	OracleNvarchar2
	OracleClob
	OracleNclob
	OracleNumber
	OracleFloat
	OracleDate
	OracleTimestamp
	OracleTimestampWithTimeZone
	OracleTimestampWithLocalTimeZone
	OracleIntervalYearToMonth
	OracleIntervalDayToSecond
	OracleBlob
	OracleRaw
	OracleLong
	OracleRowid
	OracleUrowid
)

// Stringメソッドは、OracleDataTypeの値を対応するOracleのデータ型名に変換します。
func (dt OracleDataType) Refer() string {
	return [...]string{
		"CHAR",
		"VARCHAR2",
		"NCHAR",
		"NVARCHAR2",
		"CLOB",
		"NCLOB",
		"NUMBER",
		"FLOAT",
		"DATE",
		"TIMESTAMP",
		"TIMESTAMP WITH TIME ZONE",
		"TIMESTAMP WITH LOCAL TIME ZONE",
		"INTERVAL YEAR TO MONTH",
		"INTERVAL DAY TO SECOND",
		"BLOB",
		"RAW",
		"LONG",
		"ROWID",
		"UROWID",
	}[dt]
}

// SQLServerDataTypeは、GORMのtypeタグで指定するSQL Serverのデータ型を表す列挙型です。
type SQLServerDataType int

const (
	SQLServerBigInt SQLServerDataType = iota
	SQLServerBinary
	SQLServerBit
	SQLServerChar
	SQLServerDate
	SQLServerDateTime
	SQLServerDateTime2
	SQLServerDateTimeOffset
	SQLServerDecimal
	SQLServerFloat
	SQLServerImage
	SQLServerInt
	SQLServerMoney
	SQLServerNChar
	SQLServerNText
	SQLServerNumeric
	SQLServerNVarChar
	SQLServerReal
	SQLServerSmallDateTime
	SQLServerSmallInt
	SQLServerSmallMoney
	SQLServerText
	SQLServerTime
	SQLServerTimestamp
	SQLServerTinyInt
	SQLServerUniqueIdentifier
	SQLServerVarBinary
	SQLServerVarChar
	SQLServerXML
)

// Stringメソッドは、SQLServerDataTypeの値を文字列に変換します。
func (dt SQLServerDataType) Refer() string {
	return [...]string{
		"bigint",
		"binary",
		"bit",
		"char",
		"date",
		"datetime",
		"datetime2",
		"datetimeoffset",
		"decimal",
		"float",
		"image",
		"int",
		"money",
		"nchar",
		"ntext",
		"numeric",
		"nvarchar",
		"real",
		"smalldatetime",
		"smallint",
		"smallmoney",
		"text",
		"time",
		"timestamp",
		"tinyint",
		"uniqueidentifier",
		"varbinary",
		"varchar",
		"xml",
	}[dt]
}
