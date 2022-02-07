package fluff

var vbox = new(vboxProvider)

type vboxProvider struct {
	vmrun        string
	vdiskmanager string
	mkisofs      string
	qemuimg      string
}

func (p *vboxProvider) findExecutables() bool {
	// TODO implement
	return false
}

func (p *vboxProvider) detect() bool              { return p.findExecutables() }
func (p *vboxProvider) create(i instance) error   { return nil }
func (p *vboxProvider) destroy(i instance) error  { return nil }
func (p *vboxProvider) start(i instance) error    { return nil }
func (p *vboxProvider) stop(i instance) error     { return nil }
func (p *vboxProvider) snapshot(i instance) error { return nil }
func (p *vboxProvider) list() error {
	return nil
}
