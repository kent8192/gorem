package parts

import (
	"gorem/body/parts/tag"
)

func (g *GoremParts) Column(value string) *GoremParts {
	column := &tag.Column{}
	column = column.InitTag(value)
	g.column = column
	return g
}

func (g *GoremParts) DataType(value tag.DataTypeInterface) *GoremParts {
	column := &tag.DataType{}
	column = column.InitTag(value)
	g.dataType = column
	return g
}

func (g *GoremParts) Size(value uint) *GoremParts {
	size := &tag.Size{}
	size = size.InitTag(value)
	g.size = size
	return g
}

func (g *GoremParts) PrimaryKey() *GoremParts {
	primaryKey := &tag.PrimaryKey{}
	primaryKey = primaryKey.InitTag(true)
	g.primaryKey = primaryKey
	return g
}

func (g *GoremParts) Unique() *GoremParts {
	unique := &tag.Unique{}
	unique = unique.InitTag(true)
	g.unique = unique
	return g
}

func (g *GoremParts) DefaultValue(value string) *GoremParts {
	defaultValue := &tag.Default{}
	defaultValue = defaultValue.InitTag(value)
	g.defaultValue = defaultValue
	return g
}

func (g *GoremParts) Precision(value int) *GoremParts {
	precision := &tag.Precision{}
	precision = precision.InitTag(value)
	g.precision = precision
	return g
}

func (g *GoremParts) Scale(value uint) *GoremParts {
	scale := &tag.Scale{}
	scale = scale.InitTag(value)
	g.scale = scale
	return g
}

func (g *GoremParts) NotNull() *GoremParts {
	notNull := &tag.NotNull{}
	notNull = notNull.InitTag(true)
	g.notNull = notNull
	return g
}

func (g *GoremParts) AutoIncrement() *GoremParts {
	autoIncrement := &tag.AutoIncrement{}
	autoIncrement = autoIncrement.InitTag(true)
	g.autoIncrement = autoIncrement
	return g
}

func (g *GoremParts) AutoIncrementIncrement(value int) *GoremParts {
	autoIncrementIncrement := &tag.AutoIncrementIncrement{}
	autoIncrementIncrement = autoIncrementIncrement.InitTag(value)
	g.autoIncrementIncrement = autoIncrementIncrement
	return g
}

func (g *GoremParts) Embedded() *GoremParts {
	embedded := &tag.Embedded{}
	embedded = embedded.InitTag(true)
	g.embedded = embedded
	return g
}

func (g *GoremParts) EmbeddedPrefix(value string) *GoremParts {
	embeddedPrefix := &tag.EmbeddedPrefix{}
	embeddedPrefix = embeddedPrefix.InitTag(value)
	g.embeddedPrefix = embeddedPrefix
	return g
}

func (g *GoremParts) AutoCreateTime(value string) *GoremParts {
	autoCreateTime := &tag.AutoCreateTime{}
	autoCreateTime = autoCreateTime.InitTag(value)
	g.autoCreateTime = autoCreateTime
	return g
}

func (g *GoremParts) AutoUpdateTime(value string) *GoremParts {
	autoUpdateTime := &tag.AutoUpdateTime{}
	autoUpdateTime = autoUpdateTime.InitTag(value)
	g.autoUpdateTime = autoUpdateTime
	return g
}

func (g *GoremParts) Index(value string, sort string) *GoremParts {
	index := &tag.Index{}
	index = index.InitTag(value, sort)
	g.index = index
	return g
}

func (g *GoremParts) UniqueIndex(value string, sort string) *GoremParts {
	uniqueIndex := &tag.UniqueIndex{}
	uniqueIndex = uniqueIndex.InitTag(value, sort)
	g.uniqueIndex = uniqueIndex
	return g
}

func (g *GoremParts) Check(value string) *GoremParts {
	check := &tag.Check{}
	check = check.InitTag(value)
	g.check = check
	return g
}

func (g *GoremParts) WritePermission(value string) *GoremParts {
	writePermission := &tag.WritePermission{}
	writePermission = writePermission.InitTag(value)
	g.writePermission = writePermission
	return g
}

func (g *GoremParts) ReadPermission(value string) *GoremParts {
	readPermission := &tag.ReadPermission{}
	readPermission = readPermission.InitTag(value)
	g.readPermission = readPermission
	return g
}

func (g *GoremParts) DisregardField(value string) *GoremParts {
	disregardField := &tag.DisregardField{}
	disregardField = disregardField.InitTag(value)
	g.disregardField = disregardField
	return g
}

func (g *GoremParts) Comment(value string) *GoremParts {
	comment := &tag.Comment{}
	comment = comment.InitTag(value)
	g.comment = comment
	return g
}

func (g *GoremParts) ForeignKey(value string) *GoremParts {
	foreignKey := &tag.ForeignKey{}
	foreignKey = foreignKey.InitTag(value)
	g.foreignKey = foreignKey
	return g
}

func (g *GoremParts) References(value string) *GoremParts {
	references := &tag.References{}
	references = references.InitTag(value)
	g.references = references
	return g
}

func (g *GoremParts) Many2Many(value string) *GoremParts {
	many2many := &tag.Many2Many{}
	many2many = many2many.InitTag(value)
	g.many2many = many2many
	return g
}

func (g *GoremParts) JoinForeignKey(value string) *GoremParts {
	joinForeignKey := &tag.JoinForeignKey{}
	joinForeignKey = joinForeignKey.InitTag(value)
	g.joinForeignKey = joinForeignKey
	return g
}

func (g *GoremParts) AssociationForeignKey(value string) *GoremParts {
	associationForeignKey := &tag.AssociationForeignKey{}
	associationForeignKey = associationForeignKey.InitTag(value)
	g.associationForeignKey = associationForeignKey
	return g
}

func (g *GoremParts) Polymorphic(value string) *GoremParts {
	polymorphic := &tag.Polymorphic{}
	polymorphic = polymorphic.InitTag(value)
	g.polymorphic = polymorphic
	return g
}

func (g *GoremParts) PolymorphicValue(value string) *GoremParts {
	polymorphicValue := &tag.PolymorphicValue{}
	polymorphicValue = polymorphicValue.InitTag(value)
	g.polymorphicValue = polymorphicValue
	return g
}

func (g *GoremParts) Constraint(onUpdate, onDelete string) *GoremParts {
	constraint := &tag.Constraint{}
	constraint = constraint.InitTag(onUpdate, onDelete)
	g.constraint = constraint
	return g
}

func (g *GoremParts) AssociationAutoupdate(value bool) *GoremParts {
	associationAutoupdate := &tag.AssociationAutoupdate{}
	associationAutoupdate = associationAutoupdate.InitTag(value)
	g.association_autoupdate = associationAutoupdate
	return g
}

func (g *GoremParts) AssociationAutocreate(value bool) *GoremParts {
	associationAutocreate := &tag.AssociationAutocreate{}
	associationAutocreate = associationAutocreate.InitTag(value)
	g.association_autocreate = associationAutocreate
	return g
}

func (g *GoremParts) AssociationSaveReference(value bool) *GoremParts {
	associationSaveReference := &tag.AssociationSaveReference{}
	associationSaveReference = associationSaveReference.InitTag(value)
	g.association_save_reference = associationSaveReference
	return g
}

func (g *GoremParts) Preload(value bool) *GoremParts {
	preload := &tag.Preload{}
	preload = preload.InitTag(value)
	g.preload = preload
	return g
}
