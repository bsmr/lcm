package label

type Label struct {
	identifier string
	text       string
}

func New(id, t string, options ...interface{}) (*Label, error) {
	lbl := &Label{
		identifier: id,
		text:       t,
	}

	return lbl, nil
}
