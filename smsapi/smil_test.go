package smsapi

import (
	"strings"
	"testing"
)

var (
	smilWithTextMediaObject = `
		<smil>
			<head>
				<layout>
					<region id='text0' height='100%' width='100%'/>
				</layout>
			</head>
	
			<body>
				<seq>
					<text src='some-text-uri' region='text0'/>
				</seq>
			</body>
		</smil>
	`

	smilWithImageMediaObject = `
		<smil>
			<head>
				<layout>
					<region id='img0' height='100%' width='100%'/>
				</layout>
			</head>
	
			<body>
				<seq>
					<img src='some-image-uri' region='img0'/>
				</seq>
			</body>
		</smil>
	`

	smilWithVideoMediaObject = `
		<smil>
			<head>
				<layout>
					<region id='video0' height='100%' width='100%'/>
				</layout>
			</head>
	
			<body>
				<seq>
					<video src='some-video-uri' region='video0'/>
				</seq>
			</body>
		</smil>
	`

	combinedSmil = `
		<smil>
			<head>
				<layout>
					<region id='video0' height='100%' width='100%'/>
					<region id='img1' height='100%' width='100%'/>
				</layout>
			</head>
	
			<body>
				<seq>
					<video src='some-video-uri' region='video0'/>
					<img src='some-image-uri' region='img1'/>
				</seq>
			</body>
		</smil>
	`
)

func TestSmilWithTextMediaObject(t *testing.T) {
	smil := NewSMIL()
	smil.AddText("some-text-uri")

	result := smil.GetMinifiedTplResult()

	expected := removeNewLinesAndTabsFromString(smilWithTextMediaObject)

	if result != expected {
		t.Errorf("SMIL templates doesnt match. Expected: %s Given: %s", expected, result)
	}
}

func TestSmilWithImageMediaObject(t *testing.T) {
	smil := NewSMIL()
	smil.AddImage("some-image-uri")

	result := smil.GetMinifiedTplResult()

	expected := removeNewLinesAndTabsFromString(smilWithImageMediaObject)

	if result != expected {
		t.Errorf("SMIL templates doesnt match. Expected: %s Given: %s", expected, result)
	}
}

func TestSmilWithVideoMediaObject(t *testing.T) {
	smil := NewSMIL()
	smil.AddVideo("some-video-uri")

	result := smil.GetMinifiedTplResult()

	expected := removeNewLinesAndTabsFromString(smilWithVideoMediaObject)

	if result != expected {
		t.Errorf("SMIL templates doesnt match. Expected: %s Given: %s", expected, result)
	}
}

func TestCombinedSmil(t *testing.T) {
	smil := NewSMIL()
	smil.AddVideo("some-video-uri")
	smil.AddImage("some-image-uri")

	result := smil.GetMinifiedTplResult()

	expected := removeNewLinesAndTabsFromString(combinedSmil)

	if result != expected {
		t.Errorf("SMIL templates doesnt match. Expected: %s Given: %s", expected, result)
	}
}

func TestMarshalSmil(t *testing.T) {
	smil := NewSMIL()
	smil.AddImage("some-image-uri")

	result, err := smil.MarshalJSON()

	if err != nil {
		t.Errorf("Can not Marshal: %s", smil)
	}

	expected := `"` + removeNewLinesAndTabsFromString(smilWithImageMediaObject) + `"`

	if string(result) != expected {
		t.Errorf("Marshal error Given: %s expected: %s", result, expected)
	}
}

func removeNewLinesAndTabsFromString(input string) string {
	return strings.NewReplacer("\n", "", "\t", "").Replace(input)
}
