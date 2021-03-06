package s3dm

type TriMesh struct {
	Xform
	tris []staticTri
}

func NewTriMesh(tris []*Tri) *TriMesh {
	tm := new(TriMesh)
	tm.ResetXform()

	for _, t := range(tris) {
		p1, p2, p3 := t.Points()
		tm.tris = append(tm.tris, staticTri{*p1, *p2, *p3})
	}
	return tm
}

// TODO: Prettify...
func (tm *TriMesh) Intersect(r *Ray) (*V3, *V3) { 
	first := float64(-1)
	var fi, fn *V3

	for _, t := range(tm.tris) {
		tt := staticTri{*tm.Mulv(&t.p1), *tm.Mulv(&t.p2), *tm.Mulv(&t.p3)}
		i, n := intersectTriangle(&tt, r)
		if i != nil && (first == -1 || i.Distance(r.O()) < first) {
			first = i.Distance(r.O())
			fi = i
			fn = n
		}
	}	
	return fi, fn	
}
