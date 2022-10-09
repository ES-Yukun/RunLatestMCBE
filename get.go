package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/exec"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"

	"golang.org/x/net/http2"
)

func main() {
	request, _ := http.NewRequest("GET", "https://www.minecraft.net/en-us/download/server/bedrock", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Accept-Language", "en-us")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Host", "www.google.com")
	client := &http.Client{}
	caCert, _ := ioutil.ReadFile("/etc/ssl/cert.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	client.Transport = &http2.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}
	response, _ := client.Do(request)
	defer response.Body.Close()
	r := regexp.MustCompile(`https://minecraft.azureedge.net/bin-linux/bedrock-server-.+\.zip`)
	b, _ := io.ReadAll(response.Body)
	newLink := r.FindString(string(b))
	exec.Command("/bin/bash", "-c", "curl -O /root/BedrockServer.zip "+newLink)
	exec.Command("/bin/bash", "-c", "unzip /root/BedrockServer.zip -d /root/unziped")
	exec.Command("/bin/bash", "-c", "cd /root/unziped && rm *.properties *.json *.zip")
	exec.Command("/bin/bash", "-c", "mkdir /root/minecraft /root/buckup")
	exec.Command("/bin/bash","-c")
	exec.Command("/bin/bash", "-c", "mv -f /root/unziped/* /root/minecraft")
	exec.Command("/bin/bash", "-c", "chmod +x /root/minecraft/bedrock_server")
	if _, err := os.Stat("/etc/systemd/system/minecraft.service"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			exec.Command(
				"/bin/bash", "-c", "cat << EOF > /root/buckup/buckup.sh
				while true;
					do sleep 21600;
					if [ ! -d './buckup' ]; then
						mkdir ./buckup;
					fi;
					cp -R ./worlds /root/buckup/$(date \"+\%s\").buckup;
					cd ./buckup;
						find ./ -mtime +2 -name \"*.buckup\" -type d | xargs rm -rf;
					cd ..;
				done
				EOF"
			)
		}
	}
	exec.Command("/bin/bash", "-c", "chmod +x /root/buckup/buckup.sh")
	exec.Command("/bin/bash", "-c", "/root/buckup/buckup.sh &")
	exec.Command("/bin/bash", "-c", "cd /root/minecraft && bedrock_server")
}
