package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gilgameshskytrooper/pausepizza/src/kitchen_web_server/utils"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type TokenList struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	AssociatedUser string    `json:"associatedUser"`
	Value          string    `json:"value"`
	Timestamp      time.Time `json:"timestamp"`
}

func (tl *TokenList) Initialize() {
	raw, err1 := ioutil.ReadFile(utils.AssetsDir() + "v5/tokens/list.json")
	if err1 != nil {
		fmt.Println("Could not open v5/tokens/list.json")
	}
	err2 := json.Unmarshal(raw, &tl)
	if err2 != nil {
		fmt.Println("Trouble unmarshalling the token list")
	}
}

func (tl *TokenList) WriteFile() {
	writeFile, err1 := json.MarshalIndent(tl, "", "\t")
	if err1 != nil {
		fmt.Println("could not Marshal tl")
	}
	err2 := ioutil.WriteFile(utils.AssetsDir()+"v5/tokens/list.json", writeFile, 0644)
	if err2 != nil {
		fmt.Println("Could not write Marshelled []byte into v5/tokens/list.json file")
	}
}

// This function creates a new Token object with the passed in username, a randomly generated alphanumeric sequence as the hash of length 40, and a expiration timestamp of when it will expire (the desired duration is passed as duration [the second argument])
func (tl *TokenList) GenerateNewToken(username string, duration time.Duration) {
	// Delete old tokens if any exist for a specific user
	tl.removePreviousEntryByUsername(username)

	t := Token{AssociatedUser: username, Value: generateRandomHash(), Timestamp: time.Now().Add(duration)}

	tl.Tokens = append(tl.Tokens, t)
	tl.WriteFile()
}

// DeleteToken() will find any tokens that have expired, and delete them from the list
func DeleteOldTokens() {
	tl := TokenList{}
	tl.Initialize()
	for i, token := range tl.Tokens {
		if token.Timestamp.Sub(time.Now()) < 0 {
			tl.Tokens = append(tl.Tokens[:i], tl.Tokens[i+1:]...)
		}
	}
	tl.WriteFile()
}

// removePreviousEntryByUsername() will look for any tokens associated with a specific user, and delete it
func (tl *TokenList) removePreviousEntryByUsername(username string) {
	for i, token := range tl.Tokens {
		// fmt.Println(tl.Tokens)
		if token.AssociatedUser == username {
			tl.Tokens = append(tl.Tokens[:i], tl.Tokens[i+1:]...)
		}
	}
	tl.WriteFile()
}

// removePreviousEntryByUsername() will search for a specific token, and delete it if it exists.
func (tl *TokenList) removePreviousEntryByHash(hash string) {
	for i, token := range tl.Tokens {
		if token.Value == hash {
			tl.Tokens = append(tl.Tokens[:i], tl.Tokens[i+1:]...)
		}
	}
	tl.WriteFile()
}

// A function to generate a random 40 character string using time.Now as a seed to use as a secure token.
func generateRandomHash() string {

	b := make([]byte, 40)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	src := rand.NewSource(time.Now().UnixNano())
	for i, cache, remain := 39, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Validate if token is a valid current token or not
func (tl *TokenList) Validate(hash string) bool {
	for _, token := range tl.Tokens {
		if hash == token.Value && token.Timestamp.Sub(time.Now()) > 0 {
			return true
		}
	}
	return false
}
