package model

/******************************************************************
* Object root -> SourceCode -> []Methods -> []Method
******************************************************************/

//Source code container holding method containers
type AxSourceCode struct {
	Text        string              `xml:",chardata"`
	Declaration string              `xml:"Declaration"`
	Methods     []AxMethodContainer `xml:"Methods"`
}

// Container holding text and slice of AxMethod
type AxMethodContainer struct {
	Text    string     `xml:",chardata"`
	Methods []AxMethod `xml:"Method"`
}

type AxMethod struct {
	Text   string `xml:",chardata"`
	Name   string `xml:"Name"`
	Source string `xml:"Source"`
}
