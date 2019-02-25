package instance_test

import (
	. "github.com/cloudfoundry-incubator/bosh-backup-and-restore/instance"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/instance/fakes"
	sshfakes "github.com/cloudfoundry-incubator/bosh-backup-and-restore/ssh/fakes"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"fmt"

	"bytes"
	"io"
	"log"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JobFinderFromScripts", func() {
	var logStream *bytes.Buffer
	var logger Logger
	var jobFinder *JobFinderFromScripts
	var bbrVersion = "bbr_version"
	var instanceIdentifier InstanceIdentifier

	BeforeEach(func() {
		instanceIdentifier = InstanceIdentifier{InstanceGroupName: "identifier", InstanceId: "0", Bootstrap: true}
		logStream = bytes.NewBufferString("")
		combinedLog := log.New(io.MultiWriter(GinkgoWriter, logStream), "[instance-test] ", log.Lshortfile)
		logger = boshlog.New(boshlog.LevelDebug, combinedLog)

		jobFinder = NewJobFinder(bbrVersion, logger)
	})

	Describe("FindJobs", func() {
		var remoteRunner *sshfakes.FakeRemoteRunner
		var manifestQuerier *fakes.FakeManifestQuerier
		var jobs orchestrator.Jobs
		var jobsError error

		consulAgentReleaseName := "consul-agent-release"

		BeforeEach(func() {
			remoteRunner = new(sshfakes.FakeRemoteRunner)
			remoteRunner.FindFilesReturns([]string{
				"/var/vcap/jobs/consul_agent/bin/bbr/backup",
				"/var/vcap/jobs/consul_agent/bin/bbr/restore",
				"/var/vcap/jobs/consul_agent/bin/bbr/pre-backup-lock",
				"/var/vcap/jobs/consul_agent/bin/bbr/pre-restore-lock",
				"/var/vcap/jobs/consul_agent/bin/bbr/post-backup-unlock",
				"/var/vcap/jobs/consul_agent/bin/bbr/post-restore-unlock"},
				nil)

			manifestQuerier = new(fakes.FakeManifestQuerier)
			manifestQuerier.FindReleaseNameReturns(consulAgentReleaseName, nil)
			manifestQuerier.IsJobBackupOneRestoreAllReturns(true, nil)
		})

		JustBeforeEach(func() {
			jobs, jobsError = jobFinder.FindJobs(instanceIdentifier, remoteRunner, manifestQuerier)
		})

		It("finds the jobs", func() {
			jobs, jobsError = jobFinder.FindJobs(instanceIdentifier, remoteRunner, manifestQuerier)
			By("finding the scripts", func() {
				Expect(remoteRunner.FindFilesArgsForCall(0)).To(Equal("/var/vcap/jobs/*/bin/bbr/*"))
			})

			By("logging the scripts found", func() {
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/backup"))
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/restore"))
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/pre-backup-lock"))
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/pre-restore-lock"))
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/post-backup-unlock"))
				Expect(logStream.String()).To(ContainSubstring("identifier/0/consul_agent/post-restore-unlock"))
			})

			By("calling `FindReleaseName` with the right arguments", func() {
				instanceGroupNameActual, jobNameActual := manifestQuerier.FindReleaseNameArgsForCall(0)
				Expect(instanceGroupNameActual).To(Equal(instanceIdentifier.InstanceGroupName))
				Expect(jobNameActual).To(Equal("consul_agent"))
			})

			By("calling `IsJobBackupOneRestoreAll` with the right arguments", func() {
				instanceGroupNameActual, jobNameActual := manifestQuerier.IsJobBackupOneRestoreAllArgsForCall(0)
				Expect(instanceGroupNameActual).To(Equal(instanceIdentifier.InstanceGroupName))
				Expect(jobNameActual).To(Equal("consul_agent"))
			})

			By("not returning an error", func() {
				Expect(jobsError).NotTo(HaveOccurred())
			})

			By("returning the list of jobs", func() {
				Expect(jobs).To(ConsistOf(
					NewJob(
						remoteRunner,
						"identifier/0",
						logger,
						consulAgentReleaseName,
						BackupAndRestoreScripts{
							"/var/vcap/jobs/consul_agent/bin/bbr/backup",
							"/var/vcap/jobs/consul_agent/bin/bbr/restore",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-backup-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-restore-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-backup-lock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-restore-lock",
						},
						Metadata{},
						true,
						true,
					)))
			})
		})

		Context("when the instance node is not the bootstrap node", func() {
			BeforeEach(func() {
				instanceIdentifier = InstanceIdentifier{InstanceGroupName: "identifier", InstanceId: "0", Bootstrap: false}
			})

			It("finds the jobs", func() {
				jobs, _ = jobFinder.FindJobs(instanceIdentifier, remoteRunner, manifestQuerier)
				Expect(jobs).To(ConsistOf(
					NewJob(
						remoteRunner,
						"identifier/0",
						logger,
						consulAgentReleaseName,
						BackupAndRestoreScripts{
							"/var/vcap/jobs/consul_agent/bin/bbr/backup",
							"/var/vcap/jobs/consul_agent/bin/bbr/restore",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-backup-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-restore-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-backup-lock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-restore-lock",
						},
						Metadata{},
						true,
						false,
					)))
			})
		})

		Context("when invalid scripts are present", func() {
			BeforeEach(func() {
				remoteRunner.FindFilesReturns([]string{"/var/vcap/jobs/consul_agent/bin/foobar"}, nil)
			})

			It("ignores them", func() {
				By("finding the scripts", func() {
					Expect(remoteRunner.FindFilesArgsForCall(0)).To(Equal("/var/vcap/jobs/*/bin/bbr/*"))
				})

				By("not returning an error", func() {
					Expect(jobsError).NotTo(HaveOccurred())
				})

				By("returning an empty list of jobs", func() {
					Expect(jobs).To(BeEmpty())
				})
			})
		})

		Context("when scripts are missing", func() {
			BeforeEach(func() {
				remoteRunner.FindFilesReturns([]string{}, nil)
			})

			It("does not return an error", func() {
				By("not returning an error", func() {
					Expect(jobsError).NotTo(HaveOccurred())
				})

				By("returning an empty list", func() {
					Expect(jobs).To(BeEmpty())
				})
			})
		})

		Context("when find fails fails due to an error", func() {
			BeforeEach(func() {
				remoteRunner.FindFilesReturns(nil, fmt.Errorf("no! something bad has happened"))
			})

			It("fails", func() {
				By("calling find files", func() {
					Expect(remoteRunner.FindFilesCallCount()).To(Equal(1))
				})

				By("returning an error", func() {
					Expect(jobsError).To(MatchError(SatisfyAll(
						ContainSubstring("finding scripts failed on identifier/0"),
						ContainSubstring("no! something bad has happened"),
					)))
				})
			})
		})

		Context("when manifest querier fails to find the release for a job", func() {
			BeforeEach(func() {
				manifestQuerier.FindReleaseNameReturns("", fmt.Errorf("finding the release for a job failed"))
			})

			It("does not fail", func() {
				By("setting an empty release name")
				Expect(jobs[0].Release()).To(BeEmpty())

				By("issuing a warning")
				Expect(logStream.String()).To(ContainSubstring("WARN - could not find release name for job consul_agent"))
			})
		})

		Context("when manifest querier fails to find the backupOneRestore all property for a job", func() {
			BeforeEach(func() {
				manifestQuerier.IsJobBackupOneRestoreAllReturns(false, fmt.Errorf("finding the job and release failed"))
			})

			It("does not fail", func() {
				Expect(jobsError).ToNot(HaveOccurred())

				By("setting the backupOneRestoreAll to false")
				Expect(jobs).To(ConsistOf(
					NewJob(
						remoteRunner,
						"identifier/0",
						logger,
						consulAgentReleaseName,
						BackupAndRestoreScripts{
							"/var/vcap/jobs/consul_agent/bin/bbr/backup",
							"/var/vcap/jobs/consul_agent/bin/bbr/restore",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-backup-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/post-restore-unlock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-backup-lock",
							"/var/vcap/jobs/consul_agent/bin/bbr/pre-restore-lock",
						},
						Metadata{},
						false,
						true)))
			})
		})

		Context("when metadata scripts are present", func() {
			Context("when metadata is valid", func() {
				BeforeEach(func() {
					remoteRunner.FindFilesReturns([]string{"/var/vcap/jobs/consul_agent/bin/bbr/metadata"}, nil)
					remoteRunner.RunScriptWithEnvReturns(`---
backup_name: consul_backup
backup_should_be_locked_before:
- job_name: bosh
  release: bosh`, nil)
				})

				It("attaches the metadata to the corresponding jobs", func() {
					By("executing the metadata scripts passing the correct arguments", func() {
						cmd, env, _ := remoteRunner.RunScriptWithEnvArgsForCall(0)
						Expect(cmd).To(Equal("/var/vcap/jobs/consul_agent/bin/bbr/metadata"))
						Expect(env).To(Equal(map[string]string{"BBR_VERSION": bbrVersion}))
					})

					By("adding the metadata to the returned jobs", func() {
						Expect(jobs).To(ConsistOf(
							NewJob(
								remoteRunner,
								"identifier/0",
								logger,
								consulAgentReleaseName,
								BackupAndRestoreScripts{
									"/var/vcap/jobs/consul_agent/bin/bbr/metadata",
								},
								Metadata{
									BackupName: "consul_backup",
									BackupShouldBeLockedBefore: []LockBefore{{JobName: "bosh",
										Release: "bosh"}},
								},
								true,
								true),
						))
					})

					By("not returning an error", func() {
						Expect(jobsError).NotTo(HaveOccurred())
					})
				})

				Context("and the jobFinder is configured to omit releases", func() {
					BeforeEach(func() {
						jobFinder = NewJobFinderOmitMetadataReleases(bbrVersion, logger)
					})

					It("attaches the metadata to the corresponding jobs", func() {
						By("executing the metadata scripts passing the correct arguments", func() {
							cmd, env, _ := remoteRunner.RunScriptWithEnvArgsForCall(0)
							Expect(cmd).To(Equal("/var/vcap/jobs/consul_agent/bin/bbr/metadata"))
							Expect(env).To(Equal(map[string]string{"BBR_VERSION": bbrVersion}))
						})

						By("adding the metadata to the returned jobs with empty releases", func() {
							Expect(jobs).To(ConsistOf(
								NewJob(
									remoteRunner,
									"identifier/0",
									logger,
									consulAgentReleaseName,
									BackupAndRestoreScripts{
										"/var/vcap/jobs/consul_agent/bin/bbr/metadata",
									},
									Metadata{
										BackupName: "consul_backup",
										BackupShouldBeLockedBefore: []LockBefore{{JobName: "bosh",
											Release: ""}},
									},
									true,
									true),
							))
						})

						By("not returning an error", func() {
							Expect(jobsError).NotTo(HaveOccurred())
						})
					})
				})
			})

			Context("when finding the scripts fails", func() {
				BeforeEach(func() {
					remoteRunner.FindFilesReturns(nil, fmt.Errorf("ERROR"))
				})

				It("fails", func() {
					By("returning an error", func() {
						Expect(jobsError).To(MatchError(SatisfyAll(
							ContainSubstring("finding scripts failed on identifier/0"),
							ContainSubstring("ERROR"),
						)))
					})

					By("not trying to invoke the metadata scripts", func() {
						Expect(remoteRunner.RunScriptCallCount()).To(Equal(0))
					})
				})
			})

			Context("when executing a metadata script fails", func() {
				BeforeEach(func() {
					remoteRunner.FindFilesReturns([]string{"/var/vcap/jobs/consul_agent/bin/bbr/metadata"}, nil)
					remoteRunner.RunScriptWithEnvReturns("", fmt.Errorf("blah blah blah foo"))
				})

				It("printing the location of the error, and the original error message", func() {
					Expect(jobsError).To(MatchError(ContainSubstring(
						"An error occurred while running metadata script for job consul_agent on identifier/0: blah blah blah foo",
					)))
				})
			})

			Context("when a metadata script returns invalid metadata YAML", func() {
				BeforeEach(func() {
					remoteRunner.FindFilesReturns([]string{"/var/vcap/jobs/consul_agent/bin/bbr/metadata"}, nil)
					remoteRunner.RunScriptWithEnvReturns(`this metadata is missing all the keys`, nil)
				})

				It("prints the location of the error", func() {
					Expect(jobsError).To(MatchError(ContainSubstring(
						"Parsing metadata from job consul_agent on identifier/0 failed",
					)))
				})
			})
		})
	})
})
