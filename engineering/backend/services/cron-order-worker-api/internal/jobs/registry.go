package jobs

import (
	"fmt"
	"sort"
)

type Registry struct {
	jobs map[string]Job
}

func NewRegistry() *Registry {
	return &Registry{jobs: make(map[string]Job)}
}

func (r *Registry) Register(job Job) {
	r.jobs[job.Name()] = job
}

func (r *Registry) Get(name string) (Job, error) {
	job, exists := r.jobs[name]
	if !exists {
		return nil, fmt.Errorf("job not found: %s", name)
	}
	return job, nil
}

func (r *Registry) List() []Job {
	result := make([]Job, 0, len(r.jobs))
	for _, job := range r.jobs {
		result = append(result, job)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name() < result[j].Name()
	})
	return result
}
