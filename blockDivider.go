package blockeditor

import "github.com/dracory/hb"

// blockDivider creates a divider
func (b *editor) blockDivider() *hb.Tag {
	return hb.Div().
		Class(`BlockSeparator`).
		Style(`margin: 3px 0px;`).
		Style(`clear:both; height:1px; position: relative;`).
		Style(`display: flex; justify-content: center; align-items: center;`)
}
