package conversion

import "github.com/rwv/opencc-go/dictionary"

var T2SConfig = ConfigType{
	Name:         "Traditional Chinese to Simplified Chinese",
	Segmentation: dictionary.TSPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.TSPhrasesType, dictionary.TSCharactersType},
	},
}
