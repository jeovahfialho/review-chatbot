package workflow

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type WorkflowStep struct {
	Trigger  []string    `json:"trigger"`
	Response string      `json:"response"`
	Next     interface{} `json:"next"`
	Action   string      `json:"action"`
}

type Workflow map[string]map[string]WorkflowStep

func LoadWorkflow(filePath string) (Workflow, error) {
	var workflow Workflow
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &workflow)
	return workflow, nil
}

func GetNextStep(workflow Workflow, currentState string, message string) (WorkflowStep, string) {
	log.Printf("Current State: %s", currentState)
	log.Printf("Message: %s", message)

	currentStep, ok := workflow["default"][currentState]
	if !ok {
		log.Printf("Invalid current state: %s", currentState)
		return WorkflowStep{}, "end"
	}

	log.Printf("Current Step: %+v", currentStep)

	if nextState, ok := currentStep.Next.(string); ok {
		log.Printf("Next State: %s", nextState)
		return workflow["default"][nextState], nextState
	} else if nextStates, ok := currentStep.Next.(map[string]string); ok {
		for key, state := range nextStates {
			if strings.Contains(strings.ToLower(message), key) {
				log.Printf("Matching key: %s, Next State: %s", key, state)
				return workflow["default"][state], state
			}
		}
	}

	log.Printf("No matching next state for message: %s", message)
	return WorkflowStep{}, "end"
}
