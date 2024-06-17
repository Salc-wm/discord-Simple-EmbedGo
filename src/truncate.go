package embed

func updateEmbed(content *string, limit int) {
	if len(*content) > limit {
		oldContent := *content
		*content = oldContent[:limit]
	}
}


// Truncates any embed value over the character limit.
func (e *Embed) Truncate() *Embed {
	return e.TruncateDescription().
		TruncateFields().
		TruncateFooter().
		TruncateTitle()
}

// Adds last field as InLine
func (e *Embed) MakeFieldInline() *Embed {
	length := len(e.Fields) - 1
	e.Fields[length].Inline = true

	return e
}

// Truncates fields that are too long
func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:EmbedLimitField]
	}

	for _, v := range e.Fields {
		updateEmbed(&v.Name, EmbedLimitFieldName)
		updateEmbed(&v.Value, EmbedLimitFieldValue)
	}

	return e
}

func (e *Embed) TruncateDescription() *Embed {
	updateEmbed(&e.Description, EmbedLimitDescription)

	return e
}

func (e *Embed) TruncateTitle() *Embed {
	updateEmbed(&e.Title, EmbedLimitTitle)

	return e
}

func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > EmbedLimitFooter {
		e.Footer.Text = e.Footer.Text[:EmbedLimitFooter]
	}

	return e
}
