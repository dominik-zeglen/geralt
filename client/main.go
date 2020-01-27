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
	uri        string
	httpClient http.Client
	reader     *bufio.Reader
	token      string
}

func (c *client) init() {
	c.reader = bufio.NewReader(os.Stdin)
	c.httpClient = http.Client{}
	c.uri = "http://localhost:" + utils.GetEnvOrPanic("PORT")
}

func (c *client) getInput() string {
	input, err := c.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.Trim(input, "\n")
}

func (c *client) login() error {
	fmt.Print("Login: ")
	email := c.getInput()
	loginBody := api.AuthRequest{
		Email: email,
	}

	loginBodyBytes, _ := json.Marshal(loginBody)
	loginBodyReader := bytes.NewReader(loginBodyBytes)

	res, err := http.Post(c.uri+"/auth", "application/json", loginBodyReader)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("User with this email does not exist. Please try again.\n")
		return fmt.Errorf("user %s does not exist", email)
	}

	loginResponseBodyBytes, readErr := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(readErr)
	}

	var responseBody api.AuthResponse
	decodeErr := json.
		NewDecoder(bytes.NewReader(loginResponseBodyBytes)).
		Decode(&responseBody)
	if decodeErr != nil {
		panic(decodeErr)
	}

	c.token = responseBody.Token
	return nil
}

func (c *client) ask() {
	fmt.Print("> ")
	sentence := c.getInput()

	replyBody := api.ReplyRequest{
		Sentence: sentence,
	}

	replyBodyBytes, _ := json.Marshal(replyBody)
	replyBodyReader := bytes.NewReader(replyBodyBytes)

	req, err := http.NewRequest(http.MethodPost, c.uri, replyBodyReader)
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

func (c *client) start() {
	c.init()

	for repeat := true; repeat; repeat = c.login() != nil {
	}

	for true {
		c.ask()
	}
}

func main() {
	c := client{}
	c.start()
}
