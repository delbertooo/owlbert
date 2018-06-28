package webhook

import (
	"encoding/json"
	"net/http"
)

func ReadGitlabHookMetaData(r *http.Request) (*HookMetaData, error) {
	decoder := json.NewDecoder(r.Body)
	var t HookMetaData
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

type HookMetaData struct {
	ObjectKind string `json:"object_kind"`
	ProjectId  int    `json:"project_id"`
}
