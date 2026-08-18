package geom

type HittableList struct {
	hittables []Hittable
}

func (l *HittableList) Add(h Hittable) {
	l.hittables = append(l.hittables, h)
}

func (l *HittableList) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	tempRec := HitRecord{}
	isHitAnything := false
	closestSoFar := tMax

	for _, obj := range l.hittables {
		if rec, ok := obj.Hit(r, tMin, closestSoFar); ok {
			closestSoFar = rec.T
			isHitAnything = true
			tempRec = rec
		}
	}

	return tempRec, isHitAnything
}
