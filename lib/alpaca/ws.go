package alpaca

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

type MessageType string

const (
	Success      MessageType = "success"
	Error        MessageType = "error"
	Subscription MessageType = "subscription"
)

type ServerMessage struct {
	Type      MessageType `json:"T"`
	Message   string      `json:"msg"`
	ErrorCode int         `json:"code"`
	Trades    []string    `json:"trades"`
	Quotes    []string    `json:"quotes"`
	Candles   []string    `json:"bars"`
}

type ActionType string

const (
	Authenticate ActionType = "auth"
	Subscribe    ActionType = "subscribe"
	Unsubscribe  ActionType = "unsubscribe"
)

type authAction struct {
	Action    ActionType `json:"action"`
	APIKey    string     `json:"key"`
	APISecret string     `json:"secret"`
}

type subscriptionAction struct {
	Action  ActionType `json:"action"`
	Trades  []string   `json:"trades"`
	Quotes  []string   `json:"quotes"`
	Candles []string   `json:"bars"`
}

func (c *Client) OpenRealtimeConn() error {
	if c.wsConn != nil {
		return errors.New("realtime connection already open")
	}

	source := "iex"
	if c.config.IsPaid {
		source = "sip"
	}
	u := url.URL{Scheme: "wss", Host: "stream.data.alpaca.markets", Path: "/v2/" + source}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	if _, err = getServerMessage(conn); err != nil {
		return err
	}
	if err = conn.WriteJSON(authAction{
		Action:    Authenticate,
		APIKey:    c.config.ApiKey,
		APISecret: c.config.ApiSecret,
	}); err != nil {
		return err
	}
	if _, err = getServerMessage(conn); err != nil {
		return err
	}
	c.wsConn = conn
	return nil
}

func (c *Client) sendSubscriptionAction(actionType ActionType, trades, quotes, candles []string) (msg ServerMessage, err error) {
	if c.wsConn == nil {
		return msg, errors.New("realtime connection must be opened first")
	}
	if err = c.wsConn.WriteJSON(subscriptionAction{
		Action:  actionType,
		Trades:  trades,
		Quotes:  quotes,
		Candles: candles,
	}); err != nil {
		return msg, err
	}
	msg, err = getServerMessage(c.wsConn)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (c *Client) Subscribe(trades, quotes, candles []string) (ServerMessage, error) {
	return c.sendSubscriptionAction(Subscribe, trades, quotes, candles)
}

func (c *Client) Unsubscribe(trades, quotes, candles []string) (ServerMessage, error) {
	return c.sendSubscriptionAction(Unsubscribe, trades, quotes, candles)
}

func getServerMessage(conn *websocket.Conn) (msg ServerMessage, err error) {
	var messages []ServerMessage
	if err = conn.ReadJSON(&messages); err != nil {
		return msg, err
	}
	msg = messages[0]
	if msg.Type == Error {
		return msg, errors.New(fmt.Sprintf("%d - %s", msg.ErrorCode, msg.Message))
	}
	return msg, nil
}
