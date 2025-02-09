package storage



type PageFile interface {
	GetDirPage()
	GetPage()
	NextPage()
	PrevPage()

	AllocatePage()
	WritePage()
	DeletePage()
} 