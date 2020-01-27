package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dominik-zeglen/geralt/api"
	"github.com/dominik-zeglen/geralt/utils"
)

type client struct {
	httpClient http.Client
	reader     *bufio.Reader
	token      string
}

func (c *client) start() {
	fmt.Print("Login: ")

	c.reader = bufio.NewReader(os.Stdin)
	email, _ := c.reader.ReadString('\n')

	loginBody := api.AuthRequest{
		Email: strings.Trim(email, "\n"),
	}

	loginBodyBytes, _ := json.Marshal(loginBody)
	loginBodyReader := bytes.NewReader(loginBodyBytes)

	res, err := http.Post("http://localhost:"+utils.GetEnvOrPanic("PORT")+"/auth", "application/json", loginBodyReader)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	loginResponseBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var responseBody api.AuthResponse
	decodeErr := json.NewDecoder(bytes.NewReader(loginResponseBodyBytes)).Decode(&responseBody)
	if decodeErr != nil {
		panic(decodeErr)
	}

	c.token = responseBody.Token
	c.httpClient = http.Client{}

	for true {
		fmt.Print("> ")

		sentence, _ := c.reader.ReadString('\n')

		replyBody := api.ReplyRequest{
			Sentence: strings.Trim(sentence, "\n"),
		}

		replyBodyBytes, _ := json.Marshal(replyBody)
		replyBodyReader := bytes.NewReader(replyBodyBytes)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+utils.GetEnvOrPanic("PORT"), replyBodyReader)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Authorization", "jwt "+c.token)
		req.Header.Set("Content-type", "application/json")
		res, err := c.httpClient.Do(req)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		loginResponseBodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseBody api.ReplyResponse
		decodeErr := json.NewDecoder(bytes.NewReader(loginResponseBodyBytes)).Decode(&responseBody)
		if decodeErr != nil {
			panic(decodeErr)
		}

		fmt.Println(responseBody.Reply)
	}
}

func main() {
	c := client{}
	c.start()
}
