package builder

import "testing"

func TestBuilder(t *testing.T) {
	d := Director{}

	var a IPackageBuilder
	a = new(PackageA)

	var b IPackageBuilder
	b = new(PackageB)

	d.ConstructPackage(a)
	p := d.GetPackage()

	t.Log(p)

	d.ConstructPackage(b)
	p = d.GetPackage()
	t.Log(p)
}
