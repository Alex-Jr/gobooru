package query_parser

import "fmt"

func LikeParserFn(i interface{}) interface{} {
	return fmt.Sprintf("%%%s%%", i)
}

type Direction int

const (
	RIGHT Direction = iota
	LEFT
	BOTH
)

// TODO: Add tests and remove old LikeParserFn
func LikeParserFnConfigurable(position Direction) func(i interface{}) interface{} {
	return func(i interface{}) interface{} {
		switch position {
		case RIGHT:
			return fmt.Sprintf("%s%%", i)
		case LEFT:
			return fmt.Sprintf("%%%s", i)
		case BOTH:
			return fmt.Sprintf("%%%s%%", i)
		default:
			return fmt.Sprintf("%%%s%%", i)
		}
	}
}
