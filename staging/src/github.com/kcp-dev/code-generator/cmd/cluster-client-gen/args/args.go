/*
Copyright 2025 The KCP Authors.
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package args

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/kcp-dev/code-generator/v3/cmd/cluster-client-gen/types"
)

type Args struct {
	// The directory for the generated results.
	OutputDir string

	// The Go import-path of the generated results.
	OutputPkg string

	// The boilerplate header for Go files.
	GoHeaderFile string

	// A sorted list of group versions to generate. For each of them the package path is found
	// in GroupVersionToInputPath.
	Groups []types.GroupVersions

	// Overrides for which types should be included in the client.
	IncludedTypesOverrides map[types.GroupVersion][]string

	// ClientsetAPIPath is the default API HTTP path for generated clients.
	ClientsetAPIPath string
	// ClientsetOnly determines if we should generate the clients for groups and
	// types along with the clientset. It's populated from command-line
	// arguments.
	ClientsetOnly bool
	// FakeClient determines if client-gen generates the fake clients.
	FakeClient bool
	// PluralExceptions specify list of exceptions used when pluralizing certain types.
	// For example 'Endpoints:Endpoints', otherwise the pluralizer will generate 'Endpointes'.
	PluralExceptions []string

	// Path to the generated Kubernetes single-cluster clientset package.
	SingleClusterClientPackage string
	// Path to the generated Kubernetes single-cluster applyconfigurations package.
	SingleClusterApplyConfigurationsPackage string
}

func New() *Args {
	return &Args{
		ClientsetAPIPath: "/apis",
		ClientsetOnly:    false,
		FakeClient:       true,
	}
}

func (args *Args) AddFlags(fs *pflag.FlagSet, inputBase string) {
	gvsBuilder := NewGroupVersionsBuilder(&args.Groups)
	fs.StringVar(&args.OutputDir, "output-dir", "",
		"the base directory under which to generate results")
	fs.StringVar(&args.OutputPkg, "output-pkg", args.OutputPkg,
		"the Go import-path of the generated results")
	fs.StringVar(&args.GoHeaderFile, "go-header-file", "",
		"the path to a file containing boilerplate header text; the string \"YEAR\" will be replaced with the current 4-digit year")
	fs.Var(NewGVPackagesValue(gvsBuilder, nil), "input",
		`group/versions that client-gen will generate clients for. At most one version per group is allowed. Specified in the format "group1/version1,group2/version2...".`)
	fs.Var(NewGVTypesValue(&args.IncludedTypesOverrides, []string{}), "included-types-overrides",
		"list of group/version/type for which client should be generated. By default, client is generated for all types which have genclient in types.go. This overrides that. For each groupVersion in this list, only the types mentioned here will be included. The default check of genclient will be used for other group versions.")
	fs.Var(NewInputBasePathValue(gvsBuilder, inputBase), "input-base",
		"base path to look for the api group.")
	fs.StringVarP(&args.ClientsetAPIPath, "clientset-api-path", "", args.ClientsetAPIPath,
		"the value of default API HTTP path, starting with / and without trailing /.")
	fs.BoolVar(&args.ClientsetOnly, "clientset-only", args.ClientsetOnly,
		"when set, client-gen only generates the clientset shell, without generating the individual typed clients")
	fs.BoolVar(&args.FakeClient, "fake-clientset", args.FakeClient,
		"when set, client-gen will generate the fake clientset that can be used in tests")
	fs.StringSliceVar(&args.PluralExceptions, "plural-exceptions", args.PluralExceptions,
		"list of comma separated plural exception definitions in Type:PluralizedType form")
	fs.StringVar(&args.SingleClusterClientPackage, "single-cluster-versioned-clientset-pkg", args.SingleClusterClientPackage,
		"package path to the generated Kubernetes single-cluster clientset package")
	fs.StringVar(&args.SingleClusterApplyConfigurationsPackage, "single-cluster-applyconfigurations-pkg", args.SingleClusterApplyConfigurationsPackage,
		"package path to the generated Kubernetes single-cluster applyconfigurations package")
}

func (args *Args) Validate() error {
	if len(args.OutputDir) == 0 {
		return fmt.Errorf("--output-dir must be specified")
	}
	if len(args.OutputPkg) == 0 {
		return fmt.Errorf("--output-pkg must be specified")
	}
	if len(args.ClientsetAPIPath) == 0 {
		return fmt.Errorf("--clientset-api-path cannot be empty")
	}
	if len(args.SingleClusterClientPackage) == 0 {
		return fmt.Errorf("--single-cluster-versioned-clientset-pkg cannot be empty")
	}
	if len(args.SingleClusterApplyConfigurationsPackage) == 0 {
		return fmt.Errorf("--single-cluster-applyconfigurations-pkg cannot be empty")
	}

	return nil
}

// GroupVersionPackages returns a map from GroupVersion to the package with the types.go.
func (args *Args) GroupVersionPackages() map[types.GroupVersion]string {
	res := map[types.GroupVersion]string{}
	for _, pkg := range args.Groups {
		for _, v := range pkg.Versions {
			res[types.GroupVersion{Group: pkg.Group, Version: v.Version}] = v.Package
		}
	}
	return res
}
