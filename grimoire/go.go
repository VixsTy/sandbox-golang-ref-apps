package grimoire

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Go is a mage namespace which will manage golang actions
type Go mg.Namespace

// Generate go code
func (Go) Generate() error {
	color.Cyan("## Generate code")
	// mg.SerialDeps(Gen.Protobuf, Gen.Mocks, Gen.Migrations, Gen.Decorators, Gen.Wire)
	mg.SerialDeps(Gen.Protobuf)
	return nil
}

// Test run all test
func (Go) Test() error {
	color.Cyan("## Running unit tests")
	sh.Run("mkdir", "-p", "test-results/junit")
	return sh.RunV("gotestsum", "--junitfile", "test-results/junit/unit-tests.xml", "--", "-short", "-race", "-cover", "./...")
}

// IntegrationTest run test tagged integration
func (Go) IntegrationTest() {
	color.Cyan("## Running integration tests")
	sh.Run("mkdir", "-p", "test-results/junit")

	runIntegrationTest("Repositories", "go.zenithar.org/spotigraph/internal/repositories/test/integration")
}

// Tidy add/remove depenedencies.
func (Go) Tidy() error {
	fmt.Println("## Cleaning go modules")
	return sh.RunV("go", "mod", "tidy", "-v")
}

// Deps install dependency tools.
func (Go) Deps() error {
	color.Cyan("## Vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}

// License check that project dont used blacklisted licenses
func (Go) License() error {
	color.Cyan("## Check license")
	return sh.RunV("wwhrd", "check")
}

// Format runs gofmt on everything
func (Go) Format() error {
	color.Cyan("## Format everything")
	args := []string{"-s", "-w"}
	args = append(args, goFiles...)
	return sh.RunV("gofumpt", args...)
}

// Import runs goimports on everything
func (Go) Import() error {
	color.Cyan("## Process imports")
	args := []string{"-w"}
	args = append(args, goSrcFiles...)
	return sh.RunV("goreturns", args...)
}

// Lint run linter.
func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("## Lint go code")
	return sh.RunV("golangci-lint", "run")
}
