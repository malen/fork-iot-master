package pool

func init() {
	err := Load()
	if err != nil {
		_ = Store()
	}
}