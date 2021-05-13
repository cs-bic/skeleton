package skeleton
import "errors"
func Generate(hasher func([]byte) ([]byte, error), parent []byte, path [][]byte) ([]byte, error) {
	if hasher == nil {
		return nil, errors.New("skeleton.Generate: The hasher is nil.")
	}
	if parent == nil {
		return nil, errors.New("skeleton.Generate: The parent is nil.")
	}
	if len(path) == 0 {
		return nil, errors.New("skeleton.Generate: The path has no entries.")
	}
	for _, value := range path {
		var issue error
		parent, issue = hasher(append(parent, value...))
		if issue != nil {
			return nil, issue
		}
	}
	return parent, nil
}
