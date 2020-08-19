package plugin

type MessageCard struct {
	Type       string
	Context    string
	ThemeColor string
	Summary    string
	Sections   []MessageCardSection
}

type MessageCardSection struct {
	ActivityTitle    string
	ActivitySubtitle string
	ActivityImage    string
	Facts            []MessageCardSectionFact
	Markdown         bool
}

type MessageCardSectionFact struct {
	Name  string
	Value string
}
