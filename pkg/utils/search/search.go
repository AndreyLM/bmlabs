package search

// Pagination - pagination
type Pagination struct {
	Limit int64 `json:"limit,omitempty"`
	Page  int64 `json:"page,omitempty"`
}

// Sort - sort for search
type Sort struct {
	Attribute string `json:"attribute,omitempty"`
	Type      string `json:"type,omitempty"`
}

// FilterOperator - filter type
type FilterOperator string

const (
	// FilterOperatorIn - in operator
	FilterOperatorIn = FilterOperator("in")
	// FilterOperatorLike - in operator
	FilterOperatorLike = FilterOperator("like")
	// FilterOperatorEq - in operator
	FilterOperatorEq = FilterOperator("eq")
	// FilterOperatorNeq - in operator
	FilterOperatorNeq = FilterOperator("neq")
	// FilterOperatorQt - in operator
	FilterOperatorQt = FilterOperator("qt")
	// FilterOperatorQte - in operator
	FilterOperatorQte = FilterOperator("qte")
	// FilterOperatorLt - in operator
	FilterOperatorLt = FilterOperator("lt")
	// FilterOperatorLte - in operator
	FilterOperatorLte = FilterOperator("lte")
	// FilterOperatorRange - in operator
	FilterOperatorRange = FilterOperator("range")
)

// FilterType - filter type
type FilterType string

const (
	// FilterTypeNumber - number
	FilterTypeNumber = FilterType("number")
	// FilterTypeFloat - float
	FilterTypeFloat = FilterType("float")
	// FilterTypeDate - date
	FilterTypeDate = FilterType("date")
	// FilterTypeText - text
	FilterTypeText = FilterType("text")
)

// Filter - filter for search query
type Filter struct {
	Name     string         `json:"name,omitempty"`
	Operator FilterOperator `json:"operator,omitempty"`
	Type     FilterType     `json:"type,omitempty"`
	Values   []string       `json:"values,omitempty"`
}

// IsValidType - check if type is valid
func (f *Filter) IsValidType() bool {
	switch f.Type {
	case FilterTypeNumber, FilterTypeDate, FilterTypeFloat, FilterTypeText:
		return true
	}

	return false
}

// IsValidOperator - checks if operator is valid
func (f *Filter) IsValidOperator() bool {
	switch f.Operator {
	case FilterOperatorEq, FilterOperatorIn, FilterOperatorLike, FilterOperatorLt, FilterOperatorLte,
		FilterOperatorNeq, FilterOperatorQt, FilterOperatorQte, FilterOperatorRange:
		return true
	}
	return false
}

// Search - search
type Search struct {
	Filters    []Filter   `json:"filters,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Sort       Sort       `json:"sort,omitempty"`
}
