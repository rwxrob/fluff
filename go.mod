module github.com/rwxrob/fluff

go 1.17

require github.com/rwxrob/cmdbox v0.5.0

require gopkg.in/yaml.v2 v2.4.0 // indirect

// TODO completely remove (or comment this before release)
// (Breaks 1.17+ installs completely otherwise.)
replace github.com/rwxrob/cmdbox => ../cmdbox
