// Package spamwatch provides a HTTP wrapper for the SpamWatch API
package spamwatch

import (
	"encoding/json"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

// Client is a basic client tyoe
type Client struct {
	// Log is used to log HTTP requests and responses.
	Log *zap.SugaredLogger

	// BaseReq used to make and read HTTP requests
	BaseReq Requester

	// SWType consists all type of SpamWatch API
	SWType
}

// NewClient creates a new client
func NewClient(la *zap.SugaredLogger, endpoint string, token string) (*Client, error) {
	client := new(Client)
	client.Log = la
	client.BaseReq = &DefaultApiReq
	DefaultApiReq.token = token

	if endpoint != "" {
		DefaultApiReq.apiUrl = endpoint
		la.Debug("ApiUrl using endpoint provided by user.")
	}

	la.Debug("New Client has been created.")
	return client, nil
}


// GetBan gets a full info of a user if he/she was banned.
// This action requires User permission.
func (c *Client) GetBan(userId int) (*Ban, error) {
	var ret *Ban
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "banlist/" + strconv.Itoa(userId), nil)
	_= json.Unmarshal(data, &ret)
	return ret, err
}


// GetBans gets a list of banned users.
// This action requires Root permission.
func (c *Client) GetBans() (*[]Ban, error) {
	var ret *[]Ban
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "banlist", nil)
	_ = json.Unmarshal(data, &ret)
	return ret, err
}

// GetBansMin gets a list of banned users.
// This action requires User permission.
func (c *Client) GetBansMin() ([]int, error) {
	var ret []int
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "banlist/all", nil)
	splitDataRet := strings.Split(string(data), "\n")
	for _, val:= range splitDataRet {
		convSplitDataRet, _ := strconv.Atoi(val)
		_ = append(ret, convSplitDataRet)
	}

	return ret, err
}


// AddBan adds a ban to the list.
// This action requires Admin permission.
func (c *Client) AddBan(input Ban) (bool, error) {
	data, _ := json.Marshal(input)
	_, err := c.BaseReq.MakeRequest(c.Log, "POST", "banlist", data)
	if err != nil {return false, err}
	return true , err
}

// DeleteBan deletes a ban.
// This action requires Admin permission.
func (c *Client) DeleteBan(userId int) (bool, error) {
	_, err := c.BaseReq.MakeRequest(c.Log, "DELETE", "banlist/" + strconv.Itoa(userId) , nil)
	if err != nil {return false, err}
	return true , err
}

// GetSelf gets the Token that the request was made with.
// This action requires User permission.
func (c *Client) GetSelf() (*Token, error) {
	var ret *Token
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "tokens/self", nil)
	_ = json.Unmarshal(data, &ret)
	return ret, err
}

// GetToken gets a Token object.
// This action requires Root permission.
func (c *Client) GetToken(tokenId int) (*Token, error) {
	var ret *Token
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "tokens/" + strconv.Itoa(tokenId), nil)
	_ = json.Unmarshal(data, &ret)
	return ret, err
}

// GetTokens gets all tokens.
// This action requires Root permission.
func (c *Client) GetTokens() (*[]Token, error) {
	var ret *[]Token
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "tokens", nil)
	_ = json.Unmarshal(data, &ret)
	return ret, err
}

// CreateToken creates a new token.
// This action requires Admin permission.
func (c *Client) CreateToken(input Token) (bool, error) {
	data, _ := json.Marshal(input)
	_, err := c.BaseReq.MakeRequest(c.Log, "POST", "tokens", data)
	if err != nil {return false, err}
	return true , err
}

// DeleteToken retires a token.
// This action requires Root permission.
func (c *Client) DeleteToken(tokenId int) (bool, error)  {
	_, err := c.BaseReq.MakeRequest(c.Log, "DELETE", "tokens/" + strconv.Itoa(tokenId) , nil)
	if err != nil {return false, err}
	return true , err
}

// GetStats get the current statistic.
// This action requires User permission.
func (c *Client) GetStats() (*Stats, error) {
	var ret *Stats
	data, err := c.BaseReq.MakeRequest(c.Log, "GET", "stats", nil)
	err = json.Unmarshal(data, &ret)
	return ret, err
}

