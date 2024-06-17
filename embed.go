package embed

import dgo "github.com/bwmarrin/discordgo"

type Embed struct {
	*dgo.MessageEmbed
}

const extraContent = " (extended)"
const (  // Constants for message embed character limits
	EmbedLimit int = 4000

	EmbedLimitTitle int = 256
	EmbedLimitField int = 25
	EmbedLimitFooter int = 2048

	EmbedLimitFieldName int = 256
	EmbedLimitFieldValue int = 1024
	EmbedLimitDescription int = 2048
)

// Returns a new embed object
func NewEmbed() *Embed {
	return &Embed{ &dgo.MessageEmbed{} }
}

func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > EmbedLimitDescription {
		description = description[:EmbedLimitDescription]
	}

	e.Description = description
	return e
}

func (e *Embed) AddField(name, value string) *Embed {
	fields := make([]*dgo.MessageEmbedField, 0)

	if len(value) < EmbedLimitFieldValue {
		fields = append(fields, &dgo.MessageEmbedField{
			Name:  name,
			Value: value,
		})

	} else {
		limitFieldValue := EmbedLimitFieldValue
		extended := false

		for limitFieldValue = EmbedLimitFieldValue; limitFieldValue < len(value); {
			if limitFieldValue != EmbedLimitFieldValue && extended == false {
				name += extraContent
				extended = true
			}

			fieldContent := value[limitFieldValue]
			if fieldContent == []byte(" ")[0] || fieldContent == []byte("\n")[0] || fieldContent == []byte("-")[0] {
				fields = append(fields, &dgo.MessageEmbedField{
					Name:  name,
					Value: value[limitFieldValue-EmbedLimitFieldValue: limitFieldValue],
				})

			} else {
				fields = append(fields, &dgo.MessageEmbedField{
					Name:  name,
					Value: value[limitFieldValue-EmbedLimitFieldValue: limitFieldValue-1] + "-",
				})

				limitFieldValue--
			}

			if (limitFieldValue + EmbedLimitFieldValue) > len(value) {
				break
			}

			limitFieldValue += EmbedLimitFieldValue
		}

		if limitFieldValue < len(value) {
			name += extraContent

			fields = append(fields, &dgo.MessageEmbedField{
				Name:  name,
				Value: value[limitFieldValue:],
			})
		}
	}

	e.Fields = append(e.Fields, fields...)
	return e
}

// [Text] [iconURL]
func (e *Embed) SetFooter(args ...string) *Embed {
	var iconURL, text, proxyURL string

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough

	case len(args) > 1:
		iconURL = args[1]
		fallthrough

	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Footer = &dgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}

	return e
}


func getArgsEmbed(IDs [2]int, args ...string) (arg1, arg2 string) {
	if len(args) > IDs[0] {
		arg1 = args[IDs[0]]
	}

	if len(args) > IDs[1] {
		arg2 = args[IDs[1]]
	}

	return
}

func (e *Embed) defineGenericItem(IDs [2]int, args ...string) (URL, proxyURL string) {
	if len(args) == 0 {
		return "", ""
	}

	URL, proxyURL = getArgsEmbed(IDs, args...)
	return
}


func (e *Embed) SetImage(args ...string) {
	URL, proxyURL := e.defineGenericItem(
		[2]int{0, 1}, args...,
	)

	e.Image = &dgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
}

func (e *Embed) SetThumbnail(args ...string) {
	URL, proxyURL := e.defineGenericItem(
		[2]int{0, 1}, args...,
	)

	e.Thumbnail = &dgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
}

func (e *Embed) SetAuthor(args ...string) *Embed {
	URL, proxyURL := e.defineGenericItem([2]int{2, 3}, args...)
	name, iconURL := getArgsEmbed([2]int{0, 1}, args...)

	e.Author = &dgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,

		URL:          URL,
		ProxyIconURL: proxyURL,
	}

	return e
}

func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

func (e *Embed) SetColor(clr int) *Embed {
	e.Color = clr
	return e
}

// Sets all fields in the embed to be inline
func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}

	return e
}
