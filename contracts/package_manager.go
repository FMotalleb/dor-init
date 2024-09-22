// Package contracts describes how should be implemented
package contracts

type PackageManager interface {
	ListPackages() ([]Package, error)
	InstallPackage(Package) error
}
