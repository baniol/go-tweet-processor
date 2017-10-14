package api

import (
	"context"
	"net/http"

	"github.com/baniol/go-tweet-processor/web"
)

func (rh *requestHandler) countHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	count, err := rh.dbConn.CountTweets()
	if err != nil {
		return err
	}
	c := struct {
		Count int `json:"count"`
	}{count}

	web.Respond(ctx, w, c, 200)
	return nil
}

func (rh *requestHandler) authorsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	authors, err := rh.dbConn.GetAuthors()
	if err != nil {
		return err
	}

	web.Respond(ctx, w, authors, 200)
	return nil
}

func (rh *requestHandler) tagsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	tags, err := rh.dbConn.GetTags()
	// TODO: errors.Wrap ?
	// TODO: how does returning errors to user work ?
	if err != nil {
		return err
	}

	web.Respond(ctx, w, tags, 200)
	return nil
}

func (rh *requestHandler) authorTweetsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	name := params["name"]
	tweets, err := rh.dbConn.GetAuthorTweets(name)
	if err != nil {
		return err
	}

	web.Respond(ctx, w, tweets, 200)
	return nil
}
