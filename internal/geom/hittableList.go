package geom

type HittableList struct {
	hittables []Hittable
}

func (l *HittableList) Add(h Hittable) {
	l.hittables = append(l.hittables, h)
}

func (l *HittableList) Hit(r Ray, rayT Interval) (HitRecord, bool) {
	tempRec := HitRecord{}
	isHitAnything := false
	closestSoFar := rayT.Max

	for _, obj := range l.hittables {
		if rec, ok := obj.Hit(r, Interval{Min: rayT.Min, Max: closestSoFar}); ok {
			closestSoFar = rec.T
			isHitAnything = true
			tempRec = rec
		}
	}

	return tempRec, isHitAnything
}
