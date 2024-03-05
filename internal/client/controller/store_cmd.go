package controller

// StoreCmd is subcommand for data storage
type StoreCmd struct {
	Auth StoreAuthCmd `cmd:"" help:"Store credentials"`
	Text StoreTextCmd `cmd:"" help:"Store arbitrary text"`
	Bin  StoreBinCmd  `cmd:"" help:"Store binary file"`
	Card StoreCardCmd `cmd:"" help:"Store card data"`
}

// Run executes the command
func (s *StoreCmd) Run(arg *Arg) error {
	return nil
}
