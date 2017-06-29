package system

import (
	"fmt"
	"os"

	. "github.com/onsi/gomega"
)

var RedisDeployment = DeploymentWithName("redis")
var RedisWithMetadataDeployment = DeploymentWithName("redis-with-metadata")
var RedisWithMissingScriptDeployment = DeploymentWithName("redis-with-missing-script")
var AnotherRedisDeployment = DeploymentWithName("another-redis")
var JumpboxDeployment = DeploymentWithName("jumpbox")
var JumpboxInstance = JumpboxDeployment.Instance("jumpbox", "0")

func MustHaveEnv(keyname string) string {
	val := os.Getenv(keyname)
	Expect(val).NotTo(BeEmpty(), "Need "+keyname+" for the test")
	return val
}

func BackupDirWithTimestamp(deploymentName string) string {
	return fmt.Sprintf("%s_*T*Z", deploymentName)
}

func DeploymentWithName(name string) Deployment {
	return NewDeployment(name+"-"+MustHaveEnv("TEST_ENV"), "../../fixtures/"+name+".yml")
}
