package storage

type Service struct {
	ID uint64
	Name string
	Type string
}

//db.Create(&db.Service{
//Name:      "test22",
//CreatedAt: time.Now(),
//})
//
//var service db.Service
//db.Find(&service)
//
//fmt.Println(service.Name)
