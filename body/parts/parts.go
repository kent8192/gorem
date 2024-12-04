package parts

import "gorem/body/parts/tag"

type GoremParts struct {
	column                     *tag.Column
	dataType                   *tag.DataType
	size                       *tag.Size
	primaryKey                 *tag.PrimaryKey
	unique                     *tag.Unique
	defaultValue               *tag.Default
	precision                  *tag.Precision
	scale                      *tag.Scale
	notNull                    *tag.NotNull
	autoIncrement              *tag.AutoIncrement
	autoIncrementIncrement     *tag.AutoIncrementIncrement
	embedded                   *tag.Embedded
	embeddedPrefix             *tag.EmbeddedPrefix
	autoCreateTime             *tag.AutoCreateTime
	autoUpdateTime             *tag.AutoUpdateTime
	index                      *tag.Index
	uniqueIndex                *tag.UniqueIndex
	check                      *tag.Check
	writePermission            *tag.WritePermission // <-
	readPermission             *tag.ReadPermission  // ->
	disregardField             *tag.DisregardField  // -
	comment                    *tag.Comment
	foreignKey                 *tag.ForeignKey
	references                 *tag.References
	many2many                  *tag.Many2Many
	joinForeignKey             *tag.JoinForeignKey
	associationForeignKey      *tag.AssociationForeignKey
	polymorphic                *tag.Polymorphic
	polymorphicValue           *tag.PolymorphicValue
	constraint                 *tag.Constraint
	association_autoupdate     *tag.AssociationAutoupdate
	association_autocreate     *tag.AssociationAutocreate
	association_save_reference *tag.AssociationSaveReference
	preload                    *tag.Preload
}
