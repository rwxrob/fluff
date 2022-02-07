package fluff

// detectProvider detects a provider by looking at host system and
// returns a provider according to the following priority:
//
//     * VMware Workstation Pro
//     * Oracle VirtualBox
//     * QEMU/KVM/libvirt
//
func detectProvider() provider {
	switch {
	case vmware.detect():
		return vmware
		/*
			case vbox.detect():
				return vbox
			case libvirt.detect():
				return libvirt
		*/
	}
	return nil
}
