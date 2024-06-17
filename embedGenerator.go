package embed

import (
	"fmt"
	dgo "github.com/bwmarrin/discordgo"
)

func generateGenericEmbed(title, msg string, color int) *dgo.MessageEmbed {
	return NewEmbed().
		SetTitle(title).
		SetDescription(msg).
		SetColor(color).MessageEmbed
}

func NewGenericEmbed(embedTitle, embedMsg string, replacements ...interface{}) *dgo.MessageEmbed {
	return generateGenericEmbed(
		embedTitle,
		fmt.Sprintf(embedMsg, replacements...),
		0x1c1c1c,
	)
}

func NewGenericEmbedAdvanced(embedTitle, embedMsg string, embedColor int) *dgo.MessageEmbed {
	return generateGenericEmbed(
		embedTitle, embedMsg, embedColor,
	)
}


/// Erros
func NewErrorEmbed(errorTitle, errorMsg string, replacements ...interface{}) *dgo.MessageEmbed {
	return generateGenericEmbed(
		errorTitle,
		fmt.Sprintf(errorMsg, replacements...),
		0xb40000,
	)
}

func NewErrorEmbedAdvanced(errorTitle, errorMsg string, errorColor int) *dgo.MessageEmbed {
	return generateGenericEmbed(
		errorTitle, errorMsg, errorColor,
	)
}
