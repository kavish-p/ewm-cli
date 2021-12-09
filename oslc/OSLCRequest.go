package oslc

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Auth() []*http.Cookie {

	url := "https://10.168.0.74/ccm/j_security_check"
	method := "POST"

	payload := strings.NewReader("j_username=Administrator&j_password=P@ssw0rd")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PostmanRuntime/7.28.4")
	req.Header.Add("Accept", "*/*")
	// req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))
	// fmt.Println(res.Cookies())
	// fmt.Println(res.Header.Get("X-com-ibm-team-repository-web-auth-msg"))
	// test(res.Cookies())
	return res.Cookies()
}

func CreateDefect(summary string, description string) {

	cookie := Auth()

	url := "https://10.168.0.74/ccm/oslc/contexts/_clBjsSvsEeylht3RHbzFtg/workitems/defect"
	method := "POST"

	payload := `
		<rdf:RDF
			xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
			xmlns:dcterms="http://purl.org/dc/terms/"
			xmlns:rtc_ext="http://jazz.net/xmlns/prod/jazz/rtc/ext/1.0/"
			xmlns:oslc="http://open-services.net/ns/core#"
			xmlns:acp="http://jazz.net/ns/acp#"
			xmlns:oslc_cm="http://open-services.net/ns/cm#"
			xmlns:oslc_cmx="http://open-services.net/ns/cm-x#"
			xmlns:oslc_pl="http://open-services.net/ns/pl#"
			xmlns:acc="http://open-services.net/ns/core/acc#"
			xmlns:rtc_cm="http://jazz.net/xmlns/prod/jazz/rtc/cm/1.0/"
			xmlns:process="http://jazz.net/ns/process#" >
			<rdf:Description rdf:nodeID="A0">
				<rdf:predicate rdf:resource="http://jazz.net/xmlns/prod/jazz/rtc/cm/1.0/com.ibm.team.workitem.linktype.textualReference.textuallyReferenced"/>
				<rdf:object rdf:resource="https://10.168.0.74/ccm/resource/itemName/com.ibm.team.workitem.WorkItem/1"/>
				<rdf:type rdf:resource="http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement"/>
			</rdf:Description>
			<rdf:Description>
				<dcterms:type rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Defect</dcterms:type>
				<acc:accessContext rdf:resource="https://10.168.0.74/ccm/acclist#_clBjsSvsEeylht3RHbzFtg"/>
				<oslc_cmx:project rdf:resource="https://10.168.0.74/ccm/oslc/projectareas/_clBjsSvsEeylht3RHbzFtg"/>
				<rtc_cm:filedAgainst rdf:resource="https://10.168.0.74/ccm/resource/itemOid/com.ibm.team.workitem.Category/_eV-j8CvsEeylht3RHbzFtg"/>
				<rtc_cm:type rdf:resource="https://10.168.0.74/ccm/oslc/types/_clBjsSvsEeylht3RHbzFtg/defect"/>
				<dcterms:description rdf:parseType="Literal">` + description + `</dcterms:description>
				<rdf:type rdf:resource="http://open-services.net/ns/cm#ChangeRequest"/>
				<dcterms:subject rdf:datatype="http://www.w3.org/2001/XMLSchema#string"></dcterms:subject>
				<dcterms:title rdf:parseType="Literal">` + summary + `</dcterms:title>
			</rdf:Description>
		</rdf:RDF>
	`

	payload_reader := strings.NewReader(payload)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, payload_reader)

	for i := range cookie {
		req.AddCookie(cookie[i])
		fmt.Println(cookie[i])
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/rdf+xml")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	fmt.Println(res.Header.Get("X-com-ibm-team-repository-web-auth-msg"))
}
