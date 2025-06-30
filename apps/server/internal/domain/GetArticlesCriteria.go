package domain

type GetArticlesCriteria struct {
	KeyWord  string
	Status   *uint
	Category *uint
}
