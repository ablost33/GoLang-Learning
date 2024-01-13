package exercise_jan11

const (
	url2 = "https://reqres.in/api/users?page=2"
)

func ModifyData() error {
	/*
		data := `
			"name":"morpheus",
			"job":"sandman resident"
		`
			putOperation := func(url string) (*http.Response, error) {
				buffer := bytes.NewReader([]byte(data))
				client := http.Client{}
				req, err := http.NewRequest("PUT", url, buffer)
				if err != nil {
					return nil, err
				}
				return client.Do(req)
			}
	*/
	return nil
}
