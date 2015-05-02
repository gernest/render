package render

type TemplateData map[string]interface{}

func NewTemplateData() TemplateData {
	return make(TemplateData)
}
func (d TemplateData) Add(key string, value interface{}) {
	d[key] = value
}
func (d TemplateData) Merge(data interface{}) {
	switch t := data.(type) {
	case TemplateData:
		for k, v := range t {
			d[k] = v
		}
	case map[string]interface{}:
		for k, v := range t {
			d[k] = v
		}
	}

}
