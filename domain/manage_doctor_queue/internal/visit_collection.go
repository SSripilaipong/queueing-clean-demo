package _manage_doctor_queue

type VisitCollection struct {
	Queue map[string]VisitShortInfo `json:"queue"`
}

func (q *VisitCollection) hasVisit(visitId string) bool {
	_, exists := q.Queue[visitId]
	return exists
}

func (q *VisitCollection) add(shortInfo VisitShortInfo) {
	q.Queue[shortInfo.Id] = shortInfo
}

func (q *VisitCollection) pop(id string) VisitShortInfo {
	visit, exists := q.Queue[id]
	if exists {
		delete(q.Queue, id)
	}
	return visit
}
