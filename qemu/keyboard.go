package qemu

import (
	"github.com/mikerourke/queso/qemu/cli"
)

// Language represents the WithKeyboardLayout language to use.
type Language string

// TODO: Find out what SL and SV languages are (not in text file).

const (
	LanguageArabic            Language = "ar"
	LanguageCzech             Language = "cz"
	LanguageDanish            Language = "da"
	LanguageGerman            Language = "de"
	LanguageGermanSwitzerland Language = "de-ch"
	LanguageEnglishUK         Language = "en-gb"
	LanguageEnglishUS         Language = "en-us"
	LanguageSpanish           Language = "es"
	LanguageAmharic           Language = "et"
	LanguageFinnish           Language = "fi"
	LanguageFaroese           Language = "fo"
	LanguageFrench            Language = "fr"
	LanguageFrenchDvorak      Language = "fr-be"
	LanguageFrenchCanada      Language = "fr-ca"
	LanguageFrenchSwitzerland Language = "fr-ch"
	LanguageCroatian          Language = "hr"
	LanguageHungarian         Language = "hu"
	LanguageIcelandic         Language = "is"
	LanguageItalian           Language = "it"
	LanguageJapanese          Language = "ja"
	LanguageLithuanian        Language = "lt"
	LanguageLatvian           Language = "lv"
	LanguageMacedonian        Language = "mk"
	LanguageDutch             Language = "nl"
	LanguageDutchDvorak       Language = "nl-be"
	LanguageNorwegian         Language = "no"
	LanguagePolish            Language = "pl"
	LanguagePortuguese        Language = "pt"
	LanguagePortugueseBrazil  Language = "pt-br"
	LanguageRussian           Language = "ru"
	LanguageThai              Language = "th"
	LanguageTurkish           Language = "tr"
	LanguageSL                Language = "sl"
	LanguageSV                Language = "sv"
)

// WithKeyboardLayout specifies the keyboard layout language (for example
// LanguageFrench for French). This option is only needed where it is
// not easy to get raw PC keycodes (e.g. on Macs, with some X11 servers or with
// a VNC or curses display). You don't normally need to use it on PC/Linux or
// PC/Windows hosts.
func WithKeyboardLayout(language Language) *cli.Option {
	return cli.NewOption("k", string(language))
}
