package fluff

func detectProvider() provider {
	switch {
	case vmware.detect():
		return vmware
	case vbox.detect():
		return vbox
	}
	return nil
}
