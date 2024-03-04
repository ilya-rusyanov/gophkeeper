package controller

// StoreCmd is subcommand for data storage
type StoreCmd struct {
	Auth StoreAuthCmd `cmd:"" help:"Store credentials"`
	Text StoreTextCmd `cmd:"" help:"Store arbitrary text"`
}

// Run executes the command
func (s *StoreCmd) Run(arg *Arg) error {
	return nil
}
