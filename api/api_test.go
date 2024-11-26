package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type wantData struct {
	body   string
	status int
}

type param struct {
	key   string
	value string
}
type givenData struct {
	body   string
	params []param
	uuid   string
}

type test struct {
	name  string
	input givenData
	want  wantData
}

func (dat *givenData) getBody() string {
	if dat == nil {
		return ""
	}
	return dat.body
}

func (dat *givenData) getUUID() string {
	if dat == nil {
		return ""
	}
	return dat.uuid
}

var apiPath = "/api"

// TestFullApi run a full API test
// A valid chain of requests is used to check all the different API endpoints
func TestFullApi(t *testing.T) {
	//TODO:
	//Step 0: Setup Test
	//Step 1: input data using getraenke PUT
	//Step 2: retrieve list of all getraenke with GET, check on uuid exists
	//Step 3: patch object, add $something
	//Step 4: perform a like
	//Step 5: perform a dislike
	//Step 6: perform a superlike
	//Step 7: retrieve object, check correctness
	//Step 8: delete object
	//Step 9: check if uuid was deleted from GET
	//Step 10: Teardown Test
}

// TestPutGetraenk runs tests of the api endpoint /getranke using the PUT Method
// A Multitude of valid and invalid request bodies is passed to the backend
func TestPutGetraenk(t *testing.T) {
	router := gin.Default()

	//TODO: Setup backend

	var tests = []test{
		//TODO: add proper request bodies
		{"minimal input", givenData{body: ""}, wantData{status: 201}},
		{"fully populated", givenData{body: ""}, wantData{status: 201}},
		{"invalid name", givenData{body: ""}, wantData{status: 400}},
		{"invalid likes", givenData{body: ""}, wantData{status: 400}},
		{"invalid bics", givenData{body: ""}, wantData{status: 400}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("PUT", apiPath+"/getraenk", reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}

}

// TestGetGetraenk runs test on the /getraenk API endpoint using the GET Method
// correct output after adding a getraenk as well as correct handling of malformed GET requests
func TestGetGetraenk(t *testing.T) {
	var tests = []test{
		{"expected request", givenData{}, wantData{status: 200}},
		{"random parameters", givenData{"",
			[]param{
				{"randomKey1", "randomValue1"},
				{"adasd2123inj", "23edsad"}}, ""},
			wantData{status: 200},
		},
		{"unexpected body", givenData{body: "here be bodies"}, wantData{status: 200}},
	}

	router := gin.Default()

	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("GET", apiPath+"/getraenk", reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}

}

// TestGetGetraenkByUUID runs tests on the /getraenk/{uuid} endpoint using the GET Method
// correct object retrieval, error handling and handling of malformed requests is tested
func TestGetGetraenkByUUID(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("GET", apiPath+"/getraenk/"+tt.input.getUUID(), reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}

// TestPatchGetraenkByUUID runs tests on the /getraenk/{uuid} endpoint using the PATCH Method
// correct object updating, input validation, error handling and handling of malformed requests is tested
func TestPatchGetraenkByUUID(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("PATCH", apiPath+"/getraenk/"+tt.input.getUUID(), reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}

// TestDeleteGetraenk runs tests on the /getrank/{uuid} endpoint using the DELETE method
// correct object deletion, error handling and handling of malformed requests is tested
func TestDeleteGetraenk(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("DELETE", apiPath+"/getraenk/"+tt.input.getUUID(), reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}

// TestGetraenkDislike runs tests on the /getraenk/{uuid}/dislike endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkDislike(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("POST", apiPath+"/getraenk/"+tt.input.getUUID()+"/dislike", reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}

// TestGetraenklike runs tests on the /getraenk/{uuid}/like endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkLike(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("POST", apiPath+"/getraenk/"+tt.input.getUUID()+"/like", reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}

// TestGetraenkSuperlike runs tests on the /getraenk/{uuid}/superlike endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkSuperlike(t *testing.T) {
	var tests = []test{}
	//TODO: Add tests

	router := gin.Default()
	//TODO: Setup backend

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var reqBody io.Reader
			if tt.input.getBody() != "" {
				reqBody = strings.NewReader(tt.input.getBody())
			}
			req, _ := http.NewRequest("POST", apiPath+"/getraenk/"+tt.input.getUUID()+"/superlike", reqBody)
			//add parameters to request
			if tt.input.params != nil {
				q := req.URL.Query()
				for _, param := range tt.input.params {
					q.Add(param.key, param.value)
				}
			}
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.want.status)
			if tt.want.body != "" {
				assert.Equal(t, w.Body.String(), tt.want.body)
			}
		})
	}
}
