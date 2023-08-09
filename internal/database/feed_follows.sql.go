// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: feed_follows.sql

package rssagg

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFollowFeeds = `-- name: CreateFollowFeeds :one
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type CreateFollowFeedsParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFollowFeeds(ctx context.Context, arg CreateFollowFeedsParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, createFollowFeeds,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE id = $1 AND user_id = $2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.UserID)
	return err
}

const getFeedFollows = `-- name: GetFeedFollows :many
SELECT id, created_at, updated_at, user_id, feed_id FROM feeds_follows WHERE user_id=$1
`

func (q *Queries) GetFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedsFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedsFollow
	for rows.Next() {
		var i FeedsFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
