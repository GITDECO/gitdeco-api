package usecase

import (
	"bytes"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/svg"
	"gitdeco-api/tools"
	"html/template"
	"io/ioutil"
)

type SvgUsecase struct {
}

func NewSvgUsecase() svg.Usecase {
	return &SvgUsecase{}
}

func (du *SvgUsecase) GetBadge(text string) ([]byte, error) {
	badgeSvg, _ := ioutil.ReadFile("../resource/badge.svg")
	svg, err := template.New("svg").Parse(string(badgeSvg))
	if err != nil {
		panic(&exception.Error{Key: "SVG_PARSE_ERROR", Data: ""})
	}
	var buf bytes.Buffer
	if err := svg.Execute(&buf, struct{ Text string }{Text: text}); err != nil {
		panic(&exception.Error{Key: "SVG_FILL_ERROR", Data: ""})
	}
	return buf.Bytes(), nil
}

func (du *SvgUsecase) GetTemplate(headerName, title, projectTitle, projectSubTitle, projectDescription1, projectDescription2 string) ([]byte, error) {
	templateSvg, _ := ioutil.ReadFile("../resource/template.svg")
	svg, err := template.New("svg").Parse(string(templateSvg))
	if err != nil {
		panic(&exception.Error{Key: "SVG_PARSE_ERROR", Data: ""})
	}
	var buf bytes.Buffer
	if err := svg.Execute(&buf, struct{ HeaderName string; Title string; ProjectTitle string; ProjectSubTitle string; ProjectDescription1 string; ProjectDescription2 string; }{ HeaderName: headerName, Title: title, ProjectTitle: projectTitle, ProjectSubTitle: projectSubTitle, ProjectDescription1: projectDescription1, ProjectDescription2: projectDescription2 }); err != nil {
		panic(&exception.Error{Key: "SVG_FILL_ERROR", Data: ""})
	}
	return buf.Bytes(), nil
}

func (du *SvgUsecase) GetTemplate2(headerName, title, projectTitle, projectSubTitle, projectDescription, backgroundImage, profileImage, projectIntroImage, techStackList string) ([]byte, error) {
	template2, _ := ioutil.ReadFile("../resource/template2.svg")
	svg, err := template.New("svg").Parse(string(template2))
	if err != nil {
		panic(&exception.Error{Key: "SVG_PARSE_ERROR", Data: ""})
	}
	var buf bytes.Buffer
	if err := svg.Execute(&buf, struct{ 
			HeaderName string; 
			Title string; 
			ProjectTitle string; 
			ProjectSubTitle string; 
			ProjectDescription string; 
			BackgroundImageValue string; 
			ProfileImageValue string; 
			ProjectIntroImageValue string; 
			TechStackList string;
		}{ 
			HeaderName: headerName, 
			Title: title, 
			ProjectTitle: projectTitle, 
			ProjectSubTitle: projectSubTitle, 
			ProjectDescription: projectDescription, 
			BackgroundImageValue: tools.ImageToBase64(backgroundImage), 
			ProfileImageValue: tools.ImageToBase64(profileImage), 
			ProjectIntroImageValue: tools.ImageToBase64(projectIntroImage),
			TechStackList: techStackList,
		},
	); err != nil {
		panic(&exception.Error{Key: "SVG_FILL_ERROR", Data: err.Error()})
	}
	return buf.Bytes(), nil
}
