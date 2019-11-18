package smsapi

import (
	"bytes"
	"strings"
	"text/template"
)

const SMILTemplate = `
	<smil>
		<head>
			<layout>
				{{ range $i, $item := .Items }}
					<region id='{{ .Type }}{{ $i }}' height='100%' width='100%'/>
				{{ end }}
			</layout>
		</head>

		<body>
			<seq>
				{{ range $i, $item := .Items }}
					<{{ .Type }} src='{{ .Source }}' region='{{ .Type }}{{ $i }}'/>
				{{ end }}
			</seq>
		</body>
	</smil>
`

type MediaObjectType string

const (
	imgMediaObject   = MediaObjectType("img")
	textMediaObject  = MediaObjectType("text")
	videoMediaObject = MediaObjectType("video")
)

type MediaObject struct {
	Source string
	Type   MediaObjectType
}

func NewImgMediaObject(source string) *MediaObject {
	return &MediaObject{
		Source: source,
		Type:   imgMediaObject,
	}
}

func NewTextMediaObject(source string) *MediaObject {
	return &MediaObject{
		Source: source,
		Type:   textMediaObject,
	}
}

func NewVideoMediaObject(source string) *MediaObject {
	return &MediaObject{
		Source: source,
		Type:   videoMediaObject,
	}
}

type SMIL struct {
	tpl string

	Items []*MediaObject
}

func (s *SMIL) AddImage(image string) {
	img := NewImgMediaObject(image)

	s.Items = append(s.Items, img)
}

func (s *SMIL) AddText(text string) {
	txt := NewTextMediaObject(text)

	s.Items = append(s.Items, txt)
}

func (s *SMIL) AddVideo(video string) {
	vd := NewVideoMediaObject(video)

	s.Items = append(s.Items, vd)
}

func (s *SMIL) GetTplResult() string {
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	tmpl, err := template.New("smil").Funcs(funcMap).Parse(s.tpl)

	if err != nil {
		return ""
	}

	var buf = new(bytes.Buffer)

	err = tmpl.Execute(buf, s)

	if err != nil {
		return ""
	}

	return buf.String()
}

func (s *SMIL) GetMinifiedTplResult() string {
	compiledTpl := s.GetTplResult()

	r := strings.NewReplacer("\n", "", "\t", "")

	return r.Replace(compiledTpl)
}

func (s *SMIL) String() string {
	return s.GetTplResult()
}

func (s SMIL) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.GetMinifiedTplResult() + `"`), nil
}

func NewSMIL() *SMIL {
	return &SMIL{tpl: SMILTemplate}
}
