package anki

import "os"

var url = getEnv("LEECH_REMINDER_URL", "http://localhost:8765")
var deckName = getEnv("LEECH_REMINDER_DECK_NAME", "Mining")
var expressionFieldName = getEnv("LEECH_REMINDER_EXPRESSION_FIELD_NAME", "Expression")
var readingFieldName = getEnv("LEECH_REMINDER_READING_FIELD_NAME", "Reading")
var definitionFieldName = getEnv("LEECH_REMINDER_DEFINITION_FIELD_NAME", "Definition")

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
