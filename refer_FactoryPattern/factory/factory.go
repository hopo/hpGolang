package factory

type (
	mongodb struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	Database interface {
		PutData(string, string)
		GetData(string) string
	}

	Filesystem interface {
		CreateFile(string)
		FindFile(string) string
	}

	Factory func(string) interface{}
)

func (mdb mongodb) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (mdb mongodb) GetData(query string) string {
	mdb.database[query] = data
}

func (sql sqlite) GetData(query string) string {
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext ext4) CreateFile(path string) {
	file := file{content: "EXT4 file", name: path}
	ntfs.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}
	return ntfs.files[path]
}

func (ext ext4) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}
	return ext.files[path]
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return mongodb{
			database: make(map[string]string),
		}
	case "development":
		return sqlite{
			database: make(map[string]string),
		}
	default:
		return nil
	}
}

func FilesytemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil
	}
}