package jenkins

import	"github.com/bndr/gojenkins"

func Login() error {
jenkins := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "admin", "admin")

_, err := jenkins.Init()
if err != nil {
	panic(err)
}
return err
}
