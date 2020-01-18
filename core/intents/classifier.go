package intents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/dominik-zeglen/geralt/parser"
	"github.com/goml/gobrain"
	"github.com/goml/gobrain/persist"
)

const bagOfWordsFilename = "cache/intent-bag-of-words.json"
const classifierFilename = "cache/intent-classifier.json"

func getClass(intent int) []float64 {
	classes := make([]float64, 13)
	classes[intent] = 1

	return classes
}

func _getPredictedClass(output []float64) int {
	max := output[0]
	maxIndex := 0

	for index, value := range output {
		if max < value {
			max = value
			maxIndex = index
		}
	}

	return maxIndex
}

func _validate(cls gobrain.FeedForward, input [][][]float64) {
	predictions := make([]bool, len(input))

	for predIndex, inputLine := range input {
		output := cls.Update(inputLine[0])
		predictions[predIndex] = _getPredictedClass(output) == _getPredictedClass(inputLine[1])
	}

	sum := float64(0)
	for _, p := range predictions {
		if p {
			sum++
		}
	}

	fmt.Println(sum / float64(len(predictions)))
}

func init() {
	rand.Seed(0)
}

func (predictor *IntentPredictor) initBagOfWords(training trainingDataset) {
	saved, err := ioutil.ReadFile(bagOfWordsFilename)
	if err == nil {
		json.Unmarshal(saved, &predictor.bagOfWords)
	} else {
		predictor.bagOfWords = map[string]int{}
	}

	for _, intentData := range training {
		predictor.intents = append(predictor.intents, intentData.intent)
		if err != nil {
			for _, sentence := range intentData.sentences {
				for _, token := range parser.Transform(sentence).Tokens {
					if _, ok := predictor.bagOfWords[token.Value]; !ok {
						predictor.bagOfWords[token.Value] = len(predictor.bagOfWords)
					}
				}
			}
		}
	}

	if err != nil {
		jsonData, _ := json.Marshal(&predictor.bagOfWords)
		ioutil.WriteFile(bagOfWordsFilename, jsonData, 0644)
	}
}

func (predictor *IntentPredictor) learn(trainingData trainingDataset) {
	predictor.classifier.Init(
		len(predictor.bagOfWords),
		len(predictor.bagOfWords)/2,
		13,
	)
	input := make([][][]float64, trainingData.count())

	inputIndex := 0
	for intentIndex, intentSet := range trainingData {
		for _, sentence := range intentSet.sentences {
			parsedSentence := parser.Transform(sentence)
			input[inputIndex] = [][]float64{
				predictor.getFeatures(parsedSentence),
				getClass(intentIndex),
			}
			inputIndex++
		}
	}

	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
	predictor.classifier.Train(input, 1000, 0.2, 0.4, false)

	persist.Save(classifierFilename, predictor.classifier)
}

func (predictor IntentPredictor) getFeatures(sentence parser.ParsedSentence) []float64 {
	features := make([]float64, len(predictor.bagOfWords))
	for _, token := range sentence.Tokens {
		for word, wordIndex := range predictor.bagOfWords {
			if word == token.Value {
				features[wordIndex]++
				break
			}
		}
	}

	return features
}

func (predictor *IntentPredictor) Init() {
	trainingData := getTrainingData()

	predictor.initBagOfWords(trainingData)

	classifierLoadErr := persist.Load(classifierFilename, &predictor.classifier)
	if classifierLoadErr != nil {
		predictor.learn(trainingData)
	}
}

func (predictor IntentPredictor) GetIntent(text parser.ParsedSentence) IntentPrediction {
	prediction := predictor.classifier.Update(predictor.getFeatures(text))
	ip := IntentPrediction{}

	for intentIndex, intentProbability := range prediction {
		ip[predictor.intents[intentIndex]] = intentProbability
	}

	return ip
}
