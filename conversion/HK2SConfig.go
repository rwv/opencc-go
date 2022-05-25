package conversion

import "github.com/rwv/opencc-go/dictionary"

var HK2SConfig = ConfigType{
	Name:         "Traditional Chinese (Hong Kong standard) to Simplified Chinese",
	Segmentation: dictionary.TSPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.HKVariantsRevPhrasesType, dictionary.HKVariantsRevType},
		{dictionary.TSPhrasesType, dictionary.TSCharactersType},
	},
}
