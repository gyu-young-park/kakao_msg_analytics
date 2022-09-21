package messageparser

type Parser interface {
	Parse(string) ([][]string, error)
}
