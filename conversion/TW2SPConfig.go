package conversion

import "github.com/rwv/opencc-go/dictionary"

var TW2SPConfig = ConfigType{
	Name:         "Traditional Chinese (Taiwan standard) to Simplified Chinese (with phrases)",
	Segmentation: dictionary.TSPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.TWVariantsRevPhrasesType, dictionary.TWVariantsRevType},
		{dictionary.TWPhrasesRevType},
		{dictionary.TSPhrasesType, dictionary.TSCharactersType},
	},
}
