package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/integration/helpers/fakeservicebroker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("service-access command", func() {
	var (
		userName string
	)

	BeforeEach(func() {
		userName, _ = helpers.GetCredentials()
	})

	Describe("help", func() {
		When("--help flag is set", func() {
			It("displays command usage to output", func() {
				session := helpers.CF("service-access", "--help")
				Eventually(session).Should(Exit(0))
				Expect(session).To(Say(`NAME:`))
				Expect(session).To(Say(`\s+service-access - List service access settings`))
				Expect(session).To(Say(`USAGE:`))
				Expect(session).To(Say(`\s+cf service-access \[-b BROKER\] \[-e SERVICE\] \[-o ORG\]`))
				Expect(session).To(Say(`OPTIONS:`))
				Expect(session).To(Say(`\s+-b\s+Access for plans of a particular broker`))
				Expect(session).To(Say(`\s+-e\s+Access for service name of a particular service offering`))
				Expect(session).To(Say(`\s+-o\s+Plans accessible by a particular organization`))
				Expect(session).To(Say(`SEE ALSO:`))
				Expect(session).To(Say(`\s+disable-service-access, enable-service-access, marketplace, service-brokers`))
			})
		})
	})

	When("the environment is not setup correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(false, false, ReadOnlyOrg, "service-access")
		})
	})

	When("the environment is setup correctly", func() {
		BeforeEach(func() {
			helpers.LoginCF()
			helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
		})

		When("-b is provided with a broker name that does not exist", func() {
			It("shows an error message", func() {
				session := helpers.CF("service-access", "-b", "non-existent-broker")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say(`Getting service access for broker non-existent-broker as %s\.\.\.`, userName))
				Expect(session.Err).To(Say(`(Service broker 'non-existent-broker' not found|No service offerings found for service broker 'non-existent-broker')\.`))
			})
		})

		When("-e is provided with a service name that does not exist", func() {
			It("shows an error message", func() {
				session := helpers.CF("service-access", "-e", "non-existent-service")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say(`Getting service access for service non-existent-service as %s\.\.\.`, userName))
				Expect(session.Err).To(Say(`Service offering 'non-existent-service' not found\.`))
			})
		})

		When("-o is provided with a org name that does not exist", func() {
			It("shows an error message", func() {
				session := helpers.CF("service-access", "-o", "non-existent-org")
				Eventually(session).Should(Exit(1))
				Expect(session).To(Say(`Getting service access for organization non-existent-org as %s\.\.\.`, userName))
				Expect(session.Err).To(Say(`Organization 'non-existent-org' not found`))
			})
		})

		When("there are service offerings", func() {
			var (
				orgName   string
				spaceName string

				service     string
				servicePlan string
				broker      *fakeservicebroker.FakeServiceBroker
			)

			BeforeEach(func() {
				orgName = helpers.NewOrgName()
				spaceName = helpers.NewSpaceName()
				helpers.SetupCF(orgName, spaceName)

				broker = fakeservicebroker.New()
				broker.Services[0].Plans[1].Name = helpers.GenerateHigherName(helpers.NewPlanName, broker.Services[0].Plans[0].Name)
				broker.EnsureBrokerIsAvailable()
				service = broker.ServiceName()
				servicePlan = broker.ServicePlanName()
			})

			AfterEach(func() {
				broker.Destroy()
				helpers.QuickDeleteOrg(orgName)
			})

			It("displays all service access information", func() {
				By("showing 'none' when service access is disabled")
				session := helpers.CF("service-access")
				Eventually(session).Should(Exit(0))
				Expect(session).To(Say("Getting service access as %s...", userName))
				Expect(session).To(Say(`service\s+plan\s+access\s+org`))
				Expect(session).To(Say(`%s\s+%s\s+%s`, service, servicePlan, "none"))

				By("showing 'all' when service access is enabled globally")
				Eventually(helpers.CF("enable-service-access", service)).Should(Exit(0))

				session = helpers.CF("service-access")
				Eventually(session).Should(Exit(0))
				Expect(session).To(Say("Getting service access as %s...", userName))
				Expect(session).To(Say(`service\s+plan\s+access\s+org`))
				Expect(session).To(Say(`%s\s+%s\s+%s`, service, servicePlan, "all"))
			})

			When("some services are only accessible to certain organizations", func() {
				BeforeEach(func() {
					Eventually(helpers.CF("enable-service-access", service, "-o", orgName)).Should(Exit(0))
				})

				It("shows 'limited' access to the service", func() {
					session := helpers.CF("service-access")
					Eventually(session).Should(Exit(0))
					Expect(session).To(Say("Getting service access as %s...", userName))
					Expect(session).To(Say(`service\s+plan\s+access\s+org`))
					Expect(session).To(Say(`%s\s+%s\s+%s\s+%s`, service, servicePlan, "limited", orgName))
				})
			})

			When("multiple brokers are registered and with varying service accessibility", func() {
				var (
					otherBroker *fakeservicebroker.FakeServiceBroker

					otherOrgName string
				)

				BeforeEach(func() {
					helpers.SetupCF(orgName, spaceName)

					otherBroker = fakeservicebroker.NewAlternate()
					otherBroker.Services[0].Plans[1].Name = helpers.GenerateLowerName(helpers.NewPlanName, otherBroker.Services[0].Plans[0].Name)
					otherBroker.EnsureBrokerIsAvailable()

					otherOrgName = helpers.GenerateLowerName(helpers.NewOrgName, orgName)
					helpers.CreateOrg(otherOrgName)

					Eventually(
						helpers.CF("enable-service-access",
							service,
							"-o", otherOrgName,
							"-p", servicePlan)).Should(Exit(0))
					Eventually(helpers.CF("enable-service-access", otherBroker.Services[0].Name)).Should(Exit(0))
				})

				AfterEach(func() {
					helpers.QuickDeleteOrg(otherOrgName)
					otherBroker.Destroy()
				})

				When("the -b flag is passed", func() {
					It("shows only services from the specified broker", func() {
						session := helpers.CF("service-access", "-b", otherBroker.Name())
						Eventually(session).Should(Exit(0))
						Expect(session).To(Say("Getting service access for broker %s as %s...", otherBroker.Name(), userName))
						Expect(session).To(Say(`broker:\s+%s`, otherBroker.Name()))
						Expect(session).To(Say(`service\s+plan\s+access\s+org`))
						Expect(session).To(Say(`%s\s+%s\s+%s`, otherBroker.Services[0].Name, otherBroker.Services[0].Plans[0].Name, "all"))
						Expect(string(session.Out.Contents())).NotTo(ContainSubstring(service))
					})
				})

				When("the -e flag is passed", func() {
					It("shows only services from the specified service", func() {
						session := helpers.CF("service-access", "-e", otherBroker.Services[0].Name)
						Eventually(session).Should(Exit(0))
						Expect(session).To(Say("Getting service access for service %s as %s...", otherBroker.Services[0].Name, userName))
						Expect(session).To(Say(`broker:\s+%s`, otherBroker.Name()))
						Expect(session).To(Say(`service\s+plan\s+access\s+org`))
						Expect(session).To(Say(`%s\s+%s\s+%s`, otherBroker.Services[0].Name, otherBroker.Services[0].Plans[0].Name, "all"))
						Expect(string(session.Out.Contents())).NotTo(ContainSubstring(service))
					})
				})

				When("the -o flag is passed", func() {
					It("displays only plans accessible by the specified organization", func() {
						By("not displaying brokers that were only enabled in a different org than the provided one")
						session := helpers.CF("service-access", "-o", orgName)
						Eventually(session).Should(Exit(0))
						Expect(session).To(Say(`broker:\s+%s`, otherBroker.Name()))
						Expect(session).To(Say(`%s\s+%s\s+all`,
							otherBroker.Services[0].Name,
							otherBroker.Services[0].Plans[1].Name,
						))
						Expect(session).To(Say(`%s\s+%s\s+all`,
							otherBroker.Services[0].Name,
							otherBroker.Services[0].Plans[0].Name,
						))
						Expect(string(session.Out.Contents())).NotTo(ContainSubstring(`broker:\s+%s`, broker.Name()))

						By("displaying brokers that were enabled in the provided org")
						session = helpers.CF("service-access", "-o", otherOrgName)
						Eventually(session).Should(Exit(0))
						Expect(session).To(Say(`broker:\s+%s`, broker.Name()))
						Expect(session).To(Say(`%s\s+%s\s+limited\s+%s`,
							broker.Services[0].Name,
							broker.Services[0].Plans[0].Name,
							otherOrgName,
						))
						Expect(session).To(Say(`broker:\s+%s`, otherBroker.Name()))
						Expect(session).To(Say(`%s\s+%s\s+all`,
							otherBroker.Services[0].Name,
							otherBroker.Services[0].Plans[1].Name,
						))
						Expect(session).To(Say(`%s\s+%s\s+all`,
							otherBroker.Services[0].Name,
							otherBroker.Services[0].Plans[0].Name,
						))
					})
				})
			})
		})
	})
})
