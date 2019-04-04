package websocket

import (
	"context"
	"encoding/json"

	"golang.org/x/net/websocket"
	"golang.org/x/xerrors"
)

// ReadJSON reads a json message from c into v.
func ReadJSON(ctx context.Context, c *Conn, v interface{}) error {
	typ, r, err := c.ReadMessage(ctx)
	if err != nil {
		return xerrors.Errorf("failed to read json: %v", err)
	}

	if typ != websocket.TextFrame {
		return xerrors.Errorf("unexpected frame type for json (expected TextFrame): %v", typ)
	}

	d := json.NewDecoder(r)
	err = d.Decode(v)
	if err != nil {
		return xerrors.Errorf("failed to read json: %v", err)
	}
	return nil
}

// WriteJSON writes the json message v into c.
func WriteJSON(ctx context.Context, c *Conn, v interface{}) error {
	w := c.MessageWriter(websocket.TextFrame)
	w.SetContext(ctx)
	e := json.NewEncoder(w)
	err := e.Encode(v)
	if err != nil {
		return xerrors.Errorf("failed to write json: %v", err)
	}
	err = w.Close()
	if err != nil {
		return xerrors.Errorf("failed to write json: %v", err)
	}
	return nil
}