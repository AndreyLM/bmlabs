package user

import "github.com/andreylm/bmlabs/pkg/utils/search"

// SearchForm - search form
type SearchForm struct {
	Search search.Search `json:"search,omitempty"`
}
