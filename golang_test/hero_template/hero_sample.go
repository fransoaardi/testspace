package template

import (
	"bytes"
	"net/http"

	"github.com/shiyanhui/hero/examples/app/template"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		var userList = []string{
			"Alice",
			"Bob",
			"Tom",
		}

		// Had better use buffer pool. Hero exports `GetBuffer` and `PutBuffer` for this.
		//
		// For convenience, hero also supports `io.Writer`. For example, you can also define
		// the function to `func UserList(userList []string, w io.Writer) (int, error)`,
		// and then:
		//
		//   template.UserList(userList, w)
		//
		buffer := new(bytes.Buffer)
		template.UserList(userList, buffer)
		w.Write(buffer.Bytes())
	})

	http.ListenAndServe(":8080", nil)
}
