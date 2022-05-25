package path

import(
	"strconv"
)

var(
	Data = "dat"
	Tmpl = Data+"/tmpl"
	TmplGen = Tmpl+"/gen"
	TmplSep = Tmpl+"/sep"
	Temp = Data+"/tmp"
	Pub = Data+"/pub"
	Static = Pub+"/s"
	Database = Pub+"/db"
	Post = Pub+"/p"
	LastPostIdFile = Post+"/last"
)

func
PostById(id int) string {
	return Post+"/"+strconv.Itoa(id)
}
