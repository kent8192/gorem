package parts

import (
	"fmt"
	"strings"
)

func (p *GoremParts) Render() string {
	var tags []string

	if p.column != nil {
		tag := p.column.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.dataType != nil {
		tag := p.dataType.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.size != nil {
		tag := p.size.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.primaryKey != nil {
		tag := p.primaryKey.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.unique != nil {
		tag := p.unique.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.defaultValue != nil {
		tag := p.defaultValue.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.precision != nil {
		tag := p.precision.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.scale != nil {
		tag := p.scale.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.notNull != nil {
		tag := p.notNull.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.autoIncrement != nil {
		tag := p.autoIncrement.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.autoIncrementIncrement != nil {
		tag := p.autoIncrementIncrement.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.embedded != nil {
		tag := p.embedded.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.embeddedPrefix != nil {
		tag := p.embeddedPrefix.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.autoCreateTime != nil {
		tag := p.autoCreateTime.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.autoUpdateTime != nil {
		tag := p.autoUpdateTime.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.index != nil {
		tag := p.index.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.uniqueIndex != nil {
		tag := p.uniqueIndex.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.check != nil {
		tag := p.check.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.writePermission != nil {
		tag := p.writePermission.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.readPermission != nil {
		tag := p.readPermission.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.disregardField != nil {
		tag := p.disregardField.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.comment != nil {
		tag := p.comment.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.foreignKey != nil {
		tag := p.foreignKey.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.references != nil {
		tag := p.references.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.many2many != nil {
		tag := p.many2many.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.joinForeignKey != nil {
		tag := p.joinForeignKey.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.associationForeignKey != nil {
		tag := p.associationForeignKey.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.polymorphic != nil {
		tag := p.polymorphic.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.polymorphicValue != nil {
		tag := p.polymorphicValue.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.constraint != nil {
		tag := p.constraint.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.association_autoupdate != nil {
		tag := p.association_autoupdate.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.association_autocreate != nil {
		tag := p.association_autocreate.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.association_save_reference != nil {
		tag := p.association_save_reference.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	if p.preload != nil {
		tag := p.preload.RenderTag()
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	tag := strings.Join(tags, ";")
	tag = fmt.Sprintf(`gorm:"%s"`, tag)
	return tag
}
