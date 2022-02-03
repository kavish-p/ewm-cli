package oslc

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/spf13/viper"
)

func Auth() []*http.Cookie {

	base_url := viper.Get("base_url").(string)
	username := viper.Get("ewm_username").(string)
	password := viper.Get("ewm_password").(string)

	// fmt.Println(base_url)
	// fmt.Println(username)
	// fmt.Println(password)

	url := base_url + "/ccm/j_security_check"
	method := "POST"

	payload := strings.NewReader("j_username=" + username + "&j_password=" + password)

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

func GetContext() {
	response, err := baseGETRequest("/ccm/oslc/workitems/catalog", "GET")
	if err != nil {
		log.Fatal(err)
	}

	nameQuery, _ := QueryRDF(response, "//oslc:serviceProvider[*]/oslc:ServiceProvider[1]/dcterms:title[1]")
	contextQuery, _ := QueryRDF(response, "//oslc:serviceProvider[*]/oslc:ServiceProvider[1]/oslc:details[1]/@rdf:resource")

	for i := 0; i < len(nameQuery); i++ {
		fmt.Println(contextQuery[i].InnerText() + "\t" + nameQuery[i].InnerText())
	}
}

func GetCategory(oslc_context string) {
	response, err := baseGETRequest("/ccm/oslc/categories.xml?oslc_cm.query=rtc_cm:projectArea=\""+oslc_context+"\"", "GET")
	if err != nil {
		log.Fatal(err)
	}

	filedAgainstQuery, _ := QueryRDF(response, "//rtc_cm:Category[*]/@rdf:resource")
	filedAgainstNameQuery, _ := QueryRDF(response, "//rtc_cm:Category[*]/dc:title[1]")

	for i := 0; i < len(filedAgainstQuery); i++ {
		fmt.Println(filedAgainstQuery[i].InnerText() + "\t" + filedAgainstNameQuery[i].InnerText())
	}

}

func CreateDefect(summary string, description string) {

	cookie := Auth()

	base_url := viper.Get("base_url").(string)
	oslc_context := viper.Get("oslc_context").(string)
	filedAgainstCategory := viper.Get("filedAgainstCategory").(string)
	defectType := viper.Get("defectType").(string)

	url := base_url + "/ccm/oslc/contexts/" + oslc_context + "/workitems/defect"
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
				<rdf:object rdf:resource="` + base_url + `/ccm/resource/itemName/com.ibm.team.workitem.WorkItem/1"/>
				<rdf:type rdf:resource="http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement"/>
			</rdf:Description>
			<rdf:Description>
				<dcterms:type rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Defect</dcterms:type>
				<acc:accessContext rdf:resource="` + base_url + `/ccm/acclist#` + oslc_context + `"/>
				<oslc_cmx:project rdf:resource="` + base_url + `/ccm/oslc/projectareas/` + oslc_context + `"/>
				<rtc_cm:filedAgainst rdf:resource="` + base_url + `/ccm/resource/itemOid/com.ibm.team.workitem.Category/` + filedAgainstCategory + `"/>
				<rtc_cm:type rdf:resource="` + base_url + `/ccm/oslc/types/` + oslc_context + `/` + defectType + `"/>
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

func baseGETRequest(path string, method string) (string, error) {
	cookie := Auth()
	baseURL := viper.Get("base_url").(string)

	url := baseURL + path

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, nil)

	for i := range cookie {
		req.AddCookie(cookie[i])
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/rdf+xml")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func QueryRDF(RDF string, query string) ([]*xmlquery.Node, error) {

	doc, err := xmlquery.Parse(strings.NewReader(RDF))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	queryResult := xmlquery.Find(doc, query)
	return queryResult, nil
}
