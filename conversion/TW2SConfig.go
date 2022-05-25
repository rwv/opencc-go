package conversion

import "github.com/rwv/opencc-go/dictionary"

var TW2SConfig = ConfigType{
	Name:         "Traditional Chinese (Taiwan standard) to Simplified Chinese",
	Segmentation: dictionary.TSPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.TWVariantsRevPhrasesType, dictionary.TWVariantsRevType},
		{dictionary.TSPhrasesType, dictionary.TSCharactersType},
	},
}
