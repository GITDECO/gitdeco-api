package svg

type Usecase interface {
	GetBadge(text string) ([]byte, error)
	GetTemplate(headerName, title, projectTitle, projectSubTitle, projectDescription1, projectDescription2 string) ([]byte, error)
	GetTemplate2(headerName, title, projectTitle, projectSubTitle, projectDescription, backgroundImage, profileImage, projectIntroImage, techStackList string) ([]byte, error)
}
