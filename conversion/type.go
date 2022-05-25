package conversion

import "github.com/rwv/opencc-go/dictionary"

type ConversionType int

const (
	_ ConversionType = iota
	HK2S
	S2HK
	S2T
	S2TW
	S2TWP
	T2HK
	T2S
	T2TW
	TW2S
	TW2SP
)

type ConversionChainItem []dictionary.DictionaryType

type ConfigType struct {
	Name            string
	Segmentation    dictionary.DictionaryType
	ConversionChain []ConversionChainItem
}
