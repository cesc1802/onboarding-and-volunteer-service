package dto

type User struct {
	Id            int    `json:id`
	Role_id       int    `json:role_id`
	Department_id int    `json:department_id`
	Email         string `json:email`
	Password      string `json:password`
}

func main() {

}
