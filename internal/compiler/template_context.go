package compiler

type ContextOpts struct {
	conf   map[string]string
	prefix string
	name   string
	append bool
}

func PrepareContext(builders ...ContextOpts) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func prefixKeysOf(configMap map[string]string, prefix string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func From(confs ...map[string]string) ContextOpts {
	_ = "STUB: not implemented"
	return *new(ContextOpts)
}

func (b ContextOpts) WithKeyPrefix(prefix string) ContextOpts {
	_ = "STUB: not implemented"
	return *new(ContextOpts)
}

func (b ContextOpts) WithName(name string) ContextOpts {
	_ = "STUB: not implemented"
	return *new(ContextOpts)
}

func (b ContextOpts) AddToContext() ContextOpts {
	_ = "STUB: not implemented"
	return *new(ContextOpts)
}
