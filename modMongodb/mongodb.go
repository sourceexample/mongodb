package modMongodb

func MDB_Initialize(url string) error {
	return getSingleMongoDB().Initialize(url)
}
