package conversion

import "github.com/rwv/opencc-go/dictionary"

var S2TConfig = ConfigType{
	Name:         "Simplified Chinese to Traditional Chinese",
	Segmentation: dictionary.STPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.STPhrasesType, dictionary.STCharactersType},
	},
}
