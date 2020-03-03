package reportanalyzer

import (
	"log"
	"os"
	"os/exec"
)

func wget(url, filepath string) error {
	//run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath, "--user", "amd/paulil", "--ask-password")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

//GetBuildInfoPerASIC Get build info from Jenkins build page
func GetBuildInfoPerASIC() {

	url := "http://10.67.69.71:8080/job/amd-staging-dkms-sustaining-dGPU/job/sanity-test/job/bare-metal-pro-gfx/api/xml?tree=allBuilds[description,fullDisplayName,id,timestamp]&exclude=//allBuild[not(contains(description,navi10))]"
	filepath := "./navi10_buildinfo" + ".xml"
	if err := wget(url, filepath); err != nil {
		log.Fatal(err)
	}

}

func convertToJSON() {

	xml, err := os.Open("./navi10_buildinfo.xml")
	if err != nil {
		log.Fatal("Failed to open file ", err)
	}
}
