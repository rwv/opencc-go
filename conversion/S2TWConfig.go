package conversion

import "github.com/rwv/opencc-go/dictionary"

var S2TWConfig = ConfigType{
	Name:         "Simplified Chinese to Traditional Chinese (Taiwan standard)",
	Segmentation: dictionary.STPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.STPhrasesType, dictionary.STCharactersType},
		{dictionary.TWVariantsType},
	},
}
