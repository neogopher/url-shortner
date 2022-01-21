// Package aggregate holds aggregates that combine multiple entities into full objects.
package aggregate

import "url-shortner/pkg/domain/entity"

// URL represents a type of Link.
type URL struct {
	link *entity.Link
}
