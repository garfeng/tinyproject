package main

type LangMap map[string]string

func (lm *LangMap) tr(key string) string {
	if (*lm)[key] == "" {
		return key
	}
	return (*lm)[key]
}

var (
	langList = make(map[string]*LangMap)
)

func tr(key string) string {
	if langList[kCurrentLang] != nil {
		return langList[kCurrentLang].tr(key)
	}
	return key
}

func allLanguages() []string {
	return []string{"English", "Chinese"}
}
