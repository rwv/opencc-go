package conversion

import "github.com/rwv/opencc-go/dictionary"

var S2HKConfig = ConfigType{
	Name:         "Simplified Chinese to Traditional Chinese (Hong Kong standard)",
	Segmentation: dictionary.STPhrasesType,
	ConversionChain: []ConversionChainItem{
		{dictionary.STPhrasesType, dictionary.STCharactersType},
		{dictionary.HKVariantsPhrasesType, dictionary.HKVariantsType},
	},
}
