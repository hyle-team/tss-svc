package session

import "time"

const (
	BoundaryKeygenSession  = time.Minute
	BoundarySigningSession = 10 * time.Second
)