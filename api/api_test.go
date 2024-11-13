package api_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	t.Run("minimal input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", apiPath+"/getraenk", nil)

		//TODO: generate body, evaluate usage of testing/Fuzzing

		router.ServeHTTP(w, req)

		//TODO: assert correctness of request processing
	})

	t.Run("fully populated input", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", apiPath+"/getraenk", nil)

		//TODO: generate body, evaluate usage of testing/Fuzzing

		router.ServeHTTP(w, req)

		//TODO: assert correctness of request processing
	})

	t.Run("invalid Name", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", apiPath+"/getraenk", nil)

		//TODO: generate body, evaluate usage of testing/Fuzzing

		router.ServeHTTP(w, req)

		//TODO: assert correctness of request processing
	})

	t.Run("invalid likes", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", apiPath+"/getraenk", nil)

		//TODO: generate body, evaluate usage of testing/Fuzzing

		router.ServeHTTP(w, req)

		//TODO: assert correctness of request processing
	})

	t.Run("invalid bics", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", apiPath+"/getraenk", nil)

		//TODO: generate body, evaluate usage of testing/Fuzzing

		router.ServeHTTP(w, req)

		//TODO: assert correctness of request processing
	})

}

// TestGetGetraenk runs test on the /getraenk API endpoint using the GET Method
// correct output after adding a getraenk as well as correct handling of malformed GET requests
func TestGetGetraenk(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestGetGetraenkByUUID runs tests on the /getraenk/{uuid} endpoint using the GET Method
// correct object retrieval, error handling and handling of malformed requests is tested
func TestGetGetraenkByUUID(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestPatchGetraenkByUUID runs tests on the /getraenk/{uuid} endpoint using the PATCH Method
// correct object updating, input validation, error handling and handling of malformed requests is tested
func TestPatchGetraenkByUUID(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestDeleteGetraenk runs tests on the /getrank/{uuid} endpoint using the DELETE method
// correct object deletion, error handling and handling of malformed requests is tested
func TestDeleteGetraenk(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestGetraenkDislike runs tests on the /getraenk/{uuid}/dislike endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkDislike(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestGetraenklike runs tests on the /getraenk/{uuid}/like endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkLike(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}

// TestGetraenkSuperlike runs tests on the /getraenk/{uuid}/superlike endpoint using the POST Method
// correct object updates, input validation, error handling and handling of malformed requests is tested
func TestGetraenkSuperlike(t *testing.T) {
	router := gin.Default()

	//TODO: tests
}
