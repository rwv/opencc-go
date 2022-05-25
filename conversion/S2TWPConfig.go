package conversion

import "github.com/rwv/opencc-go/dictionary"

var S2TWPConfig = ConfigType{
	Name:         "Simplified Chinese to Traditional Chinese (Taiwan standard, with phrases)",
	Segmentation: dictionary.STPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.STPhrasesType, dictionary.STCharactersType},
		{dictionary.TWPhrasesType},
		{dictionary.TWVariantsType},
	},
}
