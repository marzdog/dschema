package model

import "encoding/xml"

// Container holding a name and slice of AxFieldGroup
type AxFieldGroupContainer struct {
	Text        string         `xml:",chardata"`
	FieldGroups []AxFieldGroup `xml:"AxTableFieldGroup"`
}

//Details of a field group value
type AxFieldGroup struct {
	Text         string `xml:",chardata"`
	Name         string `xml:"Name"`
	Fields       string `xml:"Fields"`
	AutoPopulate string `xml:"AutoPopulate"`
}

//Container holding a name and slice of AxField
type AxFieldContainer struct {
	Text   string    `xml:",chardata"`
	Fields []AxField `xml:"AxDataEntityViewField"`
}

//Details of a field value
type AxField struct {
	Text                string `xml:",chardata"`
	Xmlns               string `xml:"xmlns,attr"`
	Type                string `xml:"type,attr"`
	Name                string `xml:"Name"`
	DataField           string `xml:"DataField"`
	DataSource          string `xml:"DataSource"`
	AllowEditOnCreate   string `xml:"AllowEditOnCreate"`
	AccessModifier      string `xml:"AccessModifier"`
	Mandatory           string `xml:"Mandatory"`
	AllowEdit           string `xml:"AllowEdit"`
	ComputedFieldMethod string `xml:"ComputedFieldMethod"`
	ExtendedDataType    string `xml:"ExtendedDataType"`
	EnumType            string `xml:"EnumType"`
	CountryRegionCodes  string `xml:"CountryRegionCodes"`
	IsComputedField     string `xml:"IsComputedField"`
}

type AxKeyContainer struct {
	Text string  `xml:",chardata"`
	Keys []AxKey `xml:"AxDataEntityViewKey"`
}

type AxKey struct {
	Text   string           `xml:",chardata"`
	Name   string           `xml:"Name"`
	Fields AxFieldContainer `xml:"Fields"`
}

type AxMappingContainer struct {
	Text          string           `xml:",chardata"`
	TableMappings []AxTableMapping `xml:"AxTableMapping"`
}

type AxTableMapping struct {
	Text         string         `xml:",chardata"`
	MappingTable string         `xml:"MappingTable"`
	Connections  []AxConnection `xml:"Connections"`
}

type AxConnection struct {
	Text                   string                     `xml:",chardata"`
	TableMappingConnection []AxTableMappingConnection `xml:"AxTableMappingConnection"`
}

type AxTableMappingConnection struct {
	Text       string `xml:",chardata"`
	MapField   string `xml:"MapField"`
	MapFieldTo string `xml:"MapFieldTo"`
}

type AxDataEntity struct {
	XMLName               xml.Name              `xml:"AxDataEntityView"`
	Text                  string                `xml:",chardata"`
	I                     string                `xml:"i,attr"`
	Name                  string                `xml:"Name"`
	SourceCode            AxSourceCode          `xml:"SourceCode"`
	ConfigurationKey      string                `xml:"ConfigurationKey"`
	FormRef               string                `xml:"FormRef"`
	IsObsolete            string                `xml:"IsObsolete"`
	Label                 string                `xml:"Label"`
	IsPublic              string                `xml:"IsPublic"`
	PrimaryCompanyContext string                `xml:"PrimaryCompanyContext"`
	PrimaryKey            string                `xml:"PrimaryKey"`
	PublicCollectionName  string                `xml:"PublicCollectionName"`
	PublicEntityName      string                `xml:"PublicEntityName"`
	DeleteActions         string                `xml:"DeleteActions"`
	FieldGroups           AxFieldGroupContainer `xml:"FieldGroups"`
	Fields                AxFieldContainer      `xml:"Fields"`
	Keys                  AxKeyContainer        `xml:"Keys"`
	Mappings              string                `xml:"Mappings"`
	Ranges                string                `xml:"Ranges"`
	Relations             string                `xml:"Relations"`
	StateMachines         string                `xml:"StateMachines"`
	ViewMetadata          struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"Name"`
		SourceCode struct {
			Text    string `xml:",chardata"`
			Methods struct {
				Text   string `xml:",chardata"`
				Method struct {
					Text   string `xml:",chardata"`
					Name   string `xml:"Name"`
					Source string `xml:"Source"`
				} `xml:"Method"`
			} `xml:"Methods"`
		} `xml:"SourceCode"`
		DataSources struct {
			Text                        string `xml:",chardata"`
			AxQuerySimpleRootDataSource struct {
				Text               string `xml:",chardata"`
				Name               string `xml:"Name"`
				DynamicFields      string `xml:"DynamicFields"`
				Table              string `xml:"Table"`
				DataSources        string `xml:"DataSources"`
				DerivedDataSources string `xml:"DerivedDataSources"`
				Fields             string `xml:"Fields"`
				Ranges             string `xml:"Ranges"`
				GroupBy            string `xml:"GroupBy"`
				Having             string `xml:"Having"`
				OrderBy            string `xml:"OrderBy"`
			} `xml:"AxQuerySimpleRootDataSource"`
		} `xml:"DataSources"`
	} `xml:"ViewMetadata"`
}
