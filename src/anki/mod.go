package anki

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"leech-reminder/src/models"
	"net/http"
)

func GetLeeches() ([]models.Leech, error) {
	ids, err := getLeechNoteIds()
	if err != nil {
		return nil, err
	}

	notes, err := getNotes(ids)
	if err != nil {
		return nil, err
	}

	leeches := make([]models.Leech, len(notes))
	for i, note := range notes {
		fields := note.Fields
		expression, ok := fields[expressionFieldName]
		if !ok {
			return nil, fmt.Errorf("wrong field name: %v", expressionFieldName)
		}
		reading, ok := fields[readingFieldName]
		if !ok {
			return nil, fmt.Errorf("wrong field name: %v", readingFieldName)
		}
		definition, ok := fields[definitionFieldName]
		if !ok {
			return nil, fmt.Errorf("wrong field name: %v", definitionFieldName)
		}

		leech := models.Leech{Expression: expression.Value, Reading: reading.Value, Definition: definition.Value}
		leeches[i] = leech
	}
	return leeches, nil
}

func getLeechNoteIds() ([]uint64, error) {
	return doRequest[string, []uint64]("findNotes", map[string]string{"query": fmt.Sprintf("deck:%s tag:leech", deckName)})

}

func getNotes(ids []uint64) ([]noteSchema, error) {
	return doRequest[[]uint64, []noteSchema]("notesInfo", map[string][]uint64{"notes": ids})
}

func doRequest[T string | []uint64, V []uint64 | []noteSchema](action string, params map[string]T) (V, error) {
	jsonBytes, err := json.Marshal(requestSchema[T]{Action: action, Version: 6, Params: params})
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non 200 response from Anki. Status code: %v, response data: %v", response.StatusCode, string(responseBytes))
	}

	var parsedResponse responseSchema[V]
	json.Unmarshal(responseBytes, &parsedResponse)
	if parsedResponse.Error != nil {
		return nil, errors.New(*parsedResponse.Error)
	}
	return parsedResponse.Result, nil
}
