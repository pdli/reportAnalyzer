package reportanalyzer

import (
    "bufio"
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

    url := "http://10.67.69.71:8080/job/amd-staging-dkms-sustaining-dGPU/job/sanity-test/job/bare-metal-pro-gfx/api/xml?tree=allBuilds[description,fullDisplayName,id,timestamp]&exclude=//allBuild[not(contains(description,%22navi10%22))]"
	filepath := "./navi10_buildinfo" + ".xml"
	if err := wget(url, filepath); err != nil {
		log.Fatal(err)
	}

}

func ConvertToJson() {

    cmd := exec.Command("/bin/bash", "-c", `cat navi10_buildinfo.xml | sed 's/<\/fullDisplayName>/\n/g' | sed 's/<\/allBuild>.*//' | sed '/xml/d' | sed '/workflow/d' | tee navi10.xml`);

    stdout,err := cmd.StdoutPipe()
    if err!= nil {
        log.Fatal("Error: can't obtain atdout pipe for command", err)
        return
    }

    if err := cmd.Start(); err != nil {
        log.Fatal("Error: The command is err", err)
        return
    }

    outputBuf := bufio.NewReader(stdout)

    for {
      output, _, err := outputBuf.ReadLine()

      if err != nil{
        if err.Error() != "EOF" {
          log.Fatal("Error, %s\n", err)
        }
        return
      }

      log.Printf("%s\n", string(output))
    }
}
