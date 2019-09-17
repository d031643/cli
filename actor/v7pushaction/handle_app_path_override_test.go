package v7pushaction_test

import (
	"io/ioutil"
	"os"
	"path"  // TODO: use "path/filepath" instead

	"code.cloudfoundry.org/cli/cf/util/testhelpers/matchers"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/util/pushmanifestparser"
	. "code.cloudfoundry.org/cli/actor/v7pushaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HandleAppPathOverride", func() {
	var (
		transformedManifest pushmanifestparser.Manifest
		executeErr          error

		parsedManifest pushmanifestparser.Manifest
		flagOverrides  FlagOverrides
		err            error
	)

	BeforeEach(func() {
		flagOverrides = FlagOverrides{}
		parsedManifest = pushmanifestparser.Manifest{}
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		transformedManifest, executeErr = HandleAppPathOverride(
			parsedManifest,
			flagOverrides,
		)
	})

	When("the path flag override is set", func() {
		var relativeAppFilePath string

		BeforeEach(func() {
			file, err := ioutil.TempFile("", "")
			Expect(err).NotTo(HaveOccurred())
			relativeAppFilePath = file.Name()
			defer file.Close()

			flagOverrides = FlagOverrides{
				ProvidedAppPath: relativeAppFilePath,
			}
		})

		AfterEach(func() {
			err := os.RemoveAll(relativeAppFilePath)
			Expect(err).NotTo(HaveOccurred())
		})

		When("there are multiple apps in the manifest", func() {
			BeforeEach(func() {
				parsedManifest = pushmanifestparser.Manifest{
					Applications: []pushmanifestparser.Application{
						{},
						{},
					},
				}
			})

			It("returns an error", func() {
				Expect(executeErr).To(MatchError(translatableerror.CommandLineArgsWithMultipleAppsError{}))
			})
		})

		When("there is a single app in the manifest", func() {
			BeforeEach(func() {
				parsedManifest = pushmanifestparser.Manifest{
					Applications: []pushmanifestparser.Application{
						{
							Path: "some-path",
						},
					},
					PathToManifest: "/path/to/manifest.yml",
				}
			})

			It("overrides the path for the first app in the manifest", func() {
				Expect(transformedManifest.Applications[0].Path).To(matchers.MatchPath(relativeAppFilePath))
			})

			When("the application's path is relative and passed as a flag", func() {
				var cwd string
				var absoluteAppFilehandle *os.File
				BeforeEach(func() {
					absoluteAppFilehandle, err = ioutil.TempFile("", "")
					Expect(err).NotTo(HaveOccurred())
					defer absoluteAppFilehandle.Close()
					relativeAppFilePath = path.Join(".", absoluteAppFilehandle.Name())
					flagOverrides.ProvidedAppPath = relativeAppFilePath

					// TODO: Do NOT use Chdir! it affects ALL other threads
					// Save the current working directory so you can return to it later
					cwd, err = os.Getwd()
					Expect(err).NotTo(HaveOccurred())
					// Go to root directory before executing HandleAppPathOverride()
					err = os.Chdir("/")
					Expect(err).NotTo(HaveOccurred())
				})
				AfterEach(func() {
					err = os.Chdir(cwd)
					Expect(err).NotTo(HaveOccurred())
				})

				It("doesn't override the path for the first app in the manifest", func() {
					Expect(executeErr).NotTo(HaveOccurred())
					Expect(transformedManifest.Applications[0].Path).To(matchers.MatchPath(relativeAppFilePath))
				})
			})
		})
	})

	When("the path flag override is not set", func() {
		BeforeEach(func() {
			parsedManifest = pushmanifestparser.Manifest{
				Applications: []pushmanifestparser.Application{
					{},
				},
			}
		})

		It("does not change the app path", func() {
			Expect(executeErr).NotTo(HaveOccurred())
			Expect(transformedManifest.Applications[0].Path).To(Equal(""))
		})
	})

	When("the manifest contains an invalid path", func() {
		BeforeEach(func() {
			parsedManifest = pushmanifestparser.Manifest{
				Applications: []pushmanifestparser.Application{
					{
						Path: "some-non-existent-path",
					},
				},
			}
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(pushmanifestparser.InvalidManifestApplicationPathError{
				Path: "some-non-existent-path",
			}))
		})
	})

	When("docker info is set in the manifest", func() {
		BeforeEach(func() {
			flagOverrides.ProvidedAppPath = "/some/path"

			parsedManifest.Applications = []pushmanifestparser.Application{
				{
					Name:   "some-app",
					Docker: &pushmanifestparser.Docker{},
				},
			}
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.ArgumentManifestMismatchError{
				Arg:              "--path, -p",
				ManifestProperty: "docker",
			}))
		})
	})
})
