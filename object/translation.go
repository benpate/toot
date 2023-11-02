package object

// https://docs.joinmastodon.org/entities/Translation/
type Translation struct {
	Content                string `json:"content"`                  // HTML: The translated text of the status.
	DetectedSourceLanguage string `json:"detected_source_language"` // ISO 639 Language Code: The language of the source text, as auto-detected by the machine translation provider.
	Provider               string `json:"provider"`                 // The service that provided the machine translation.
}
