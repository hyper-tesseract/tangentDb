package env


const (
	// While opening a db, create new db if does not exist
	CreateDbIfNotExist = true

	// file permissions
	READ_WRITE = 420
	READ       = 400
	WRITE      = 020

	// Read write size 4Kb
	PAGE_SIZE  = 4096
)

func GetTangentHome() string {
	return "/home/hmuhammadazeem/tangenthome"
}


// custom types defined for the sake of abstracting
type PageId int32