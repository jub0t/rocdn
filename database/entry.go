package database

import "rblx/structs"

func New() structs.Storage {
	return structs.Storage{
		Data: []structs.Image{},
	}
}

func Insert(s *structs.Storage, i structs.Image) {
	s.Data = append(s.Data, i)
}

func Get(s *structs.Storage, id int, size int) *structs.Image {
	for i := 0; i < len(s.Data); i++ {
		r := s.Data[i]

		if r.TargetId == id && r.Size == size {
			return &r
		}
	}

	return &structs.Image{}
}

func Remove(s *structs.Storage, id int) {
	var f []structs.Image

	for i := 0; i < len(s.Data); i++ {
		r := s.Data[i]

		if r.TargetId != id {
			f = append(f, r)
		}
	}

	s.Data = f
}

func Has(s *structs.Storage, id int) bool {
	for i := 0; i < len(s.Data); i++ {
		r := s.Data[i]

		if r.TargetId == id {
			return true
		}
	}

	return false
}
