package intents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/dominik-zeglen/geralt/parser"
	"github.com/goml/gobrain"
)

var cls gobrain.FeedForward
var wordBag map[string]int

func getFeatures(sentence string) []float64 {
	features := make([]float64, len(wordBag))
	for _, token := range parser.Transform(sentence).Tokens {
		for word, wordIndex := range wordBag {
			if word == token.Value {
				features[wordIndex] = 1
				break
			}
		}
	}

	return features
}

func getClass(intent int) []float64 {
	classes := make([]float64, 13)
	classes[intent] = 1

	return classes
}

func getPredictedClass(output []float64) int {
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

func initWordBag(training trainingDataset) {
	log.Println("Init word bag")

	saved, err := ioutil.ReadFile("word-bag.json")
	if err == nil {
		json.Unmarshal(saved, &wordBag)
	} else {
		wordBag = map[string]int{}
		for _, intentData := range training {
			for _, sentence := range intentData.sentences {
				for _, token := range parser.Transform(sentence).Tokens {
					if _, ok := wordBag[token.Value]; !ok {
						wordBag[token.Value] = len(wordBag)
					}
				}
			}
		}

		jsonData, _ := json.Marshal(&wordBag)
		ioutil.WriteFile("word-bag.json", jsonData, 0644)
	}
}

func validate(cls gobrain.FeedForward, input [][][]float64) {
	predictions := make([]bool, len(input))

	for predIndex, inputLine := range input {
		output := cls.Update(inputLine[0])
		predictions[predIndex] = getPredictedClass(output) == getPredictedClass(inputLine[1])
	}

	sum := float64(0)
	for _, p := range predictions {
		if p {
			sum++
		}
	}

	fmt.Println(sum / float64(len(predictions)))
}

func initCls(training trainingDataset) {
	log.Println("Init classifier")
	cls.Init(len(wordBag), len(wordBag)/2, 13)
	input := make([][][]float64, training.count())

	inputIndex := 0
	for intentIndex, intentSet := range training {
		for _, sentence := range intentSet.sentences {
			input[inputIndex] = [][]float64{getFeatures(sentence), getClass(intentIndex)}
			inputIndex++
		}
	}

	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
	cls.Train(input, 1000, 0.2, 0.4, false)
	validate(cls, input)
}

func init() {
	rand.Seed(0)

	td := getTrainingData()

	initWordBag(td)
	initCls(td)
}

func Reply(text string) {
	pred := cls.Update(getFeatures(text))
	fmt.Printf("Decision: %s\n", getTrainingData()[getPredictedClass(pred)].intent)
	for intentIndex, response := range pred {
		fmt.Printf("%s: %0.6f\n", getTrainingData()[intentIndex].intent, response)
	}
}
