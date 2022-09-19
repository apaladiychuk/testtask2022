package model

// BasicTitle struct that represents data in csv file
type BasicTitle struct {
	Tconst         string `csv:"tconst"`
	TitleType      string `csv:"titleType"`
	PrimaryTitle   string `csv:"primaryTitle"`
	OriginalTitle  string `csv:"originalTitle"`
	IsAdult        string `csv:"isAdult"`
	StartYear      string `csv:"startYear"`
	EndYear        string `csv:"endYear"`
	RuntimeMinutes string `csv:"runtimeMinutes"`
	Genres         string `csv:"genres"`
}

func (b *BasicTitle) IsGenres(i string) bool {
	return i == "" || b.Genres == i
}

func (b *BasicTitle) IsPrimaryTitle(i string) bool {
	return i == "" || b.PrimaryTitle == i
}

func (b *BasicTitle) IsOriginalTitle(i string) bool {
	return i == "" || b.OriginalTitle == i
}
