package application

import "context"

type Repository interface {
	FixText(ctx context.Context, text string) (string, error)
}
