package terminal

type mode interface {
	prompt(hostname string) string
}

type globalMode struct{}

func (m *globalMode) prompt(hostname string) string {
    return hostname + "# "
}

type configureMode struct{}

func (m *configureMode) prompt(hostname string) string {
    return hostname + "(config)# "
}

type configureInterfaceMode struct{}

func (m *configureInterfaceMode) prompt(hostname string) string {
    return hostname + "(config-if)# "
}
